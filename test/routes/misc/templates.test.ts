import { renderTemplate } from '@/templates';
import { ENV } from '@/utils';
import { request } from '../../server';

describe('templates', () => {
  it('should render proper sms template', async () => {
    // set env vars
    await request.post('/change-env').send({
      APP_NAME: 'NHost App',
    });

    const template = 'signin-passwordless-sms';
    const message = await renderTemplate(`${template}/text`, {
      appName: ENV.APP_NAME,
      locale: 'en',
      displayName: 'John Doe',
      code: '123456',
    });

    expect(message).toEqual('123456 is your verification code for NHost App.');
  });
});
