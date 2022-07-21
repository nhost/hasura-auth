import { renderTemplate } from '@/templates';
import { request } from '../../server';

describe('templates', () => {
  it('should render proper sms template', async () => {
    // set env vars
    await request.post('/change-env').send({
      APP_NAME: 'NHost App',
    });

    const template = 'signin-passwordless-sms';
    const message = await renderTemplate(`${template}/text`, {
      locale: 'en',
      displayName: 'John Doe',
      code: '123456',
    });

    expect(message).toEqual('Your code is 123456.');
  });
});
