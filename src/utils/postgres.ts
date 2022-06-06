import axios, { AxiosInstance } from 'axios';
import { format } from '@scaleleap/pg-format';

export type RunSqlResponse =
  | {
      result_type: 'TuplesOk';
      result: string[][];
    }
  | {
      result_type: 'CommandOk';
      result: null;
    };

export type JSONValue =
  | string
  | number
  | boolean
  | { [x: string]: JSONValue }
  | Array<JSONValue>;

class PostgresClient {
  private httpClient: AxiosInstance;
  constructor() {
    console.log('in constructor');

    console.log(process.env.HASURA_GRAPHQL_GRAPHQL_URL);
    console.log(process.env.HASURA_GRAPHQL_ADMIN_SECRET);

    const baseURL = process.env.HASURA_GRAPHQL_GRAPHQL_URL?.replace(
      '/v1/graphql',
      ''
    );
    console.log({ baseURL });

    this.httpClient = axios.create({
      baseURL,
      timeout: 30000,
      headers: {
        'x-hasura-admin-secret': process.env.HASURA_GRAPHQL_ADMIN_SECRET,
      },
    });
  }

  async runSql(sql: string, args: unknown[]): Promise<RunSqlResponse> {
    let res;
    try {
      res = await this.httpClient.post('/v2/query', {
        type: 'run_sql',
        args: {
          cascade: false,
          read_only: false,
          source: 'default',
          sql: format(sql, ...args),
        },
      });
    } catch (error) {
      console.log('error');
      console.log(error);
    }
    if (!res) {
      throw Error('No response');
      console.error(res.data);
    }
    return res.data as RunSqlResponse;
  }

  async runSqlParsed(sql: string, args: unknown[]) {
    const res = await this.runSql(sql, args);

    const { columns, rows } = this.parseSelect(res);

    return { columns, rows };
  }

  parseSelect(runSqlResponse: RunSqlResponse): {
    columns: string[] | undefined;
    rows: string[][] | JSONValue[];
  } {
    if (runSqlResponse.result_type !== 'TuplesOk') {
      throw Error('Not a TuplesOk response');
    }

    // First element of result is always columns
    // .shift() removes first element of result.
    const columns = runSqlResponse.result.shift();
    let rows = runSqlResponse.result;

    if (columns && columns[0] === 'row_to_json') {
      rows = rows.map((row) => JSON.parse(row[0]));
    }

    return { columns, rows };
  }
}

const postgres = new PostgresClient();

export { postgres };
