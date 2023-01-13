const autocannon = require('autocannon');
const faker = require('faker');

const createUser = () => ({
  email: `${faker.name.firstName()}.${faker.name.lastName()}-[<id>]@${faker.internet.domainName()}`,
  password: faker.internet.password(),
});

function finishedBench(err, res) {
  if (err) {
    console.warn(err);
  }
  if (res) {
    console.log(res);
  }
  process.exit(1);
}

// * Don't run autocannon so we can profile the startup only
const startup = false;

if (startup) {
  return;
}

autocannon(
  {
    url: 'http://localhost:4000',
    //   maxConnectionRequests: 1,
    // excludeErrorStats: true,
    connections: 1000,
    duration: 10,

    initialContext: { user: createUser() },
    headers: {
      'Content-Type': 'application/json',
    },
    requests: [
      {
        method: 'POST',
        path: '/signup/email-password',
        setupRequest: (req, context) => ({
          ...req,
          body: JSON.stringify(context.user),
        }),
      },
      // {
      //   method: 'POST',
      //   path: '/signin/email-password',
      //   setupRequest: (req, context) => ({
      //     ...req,
      //     path: `/signup/email-password`,
      //     body: JSON.stringify(context.user),
      //   }),
      // },
    ],
    idReplacement: true,
  },
  finishedBench
);
