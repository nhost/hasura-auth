import { request } from './server';
import { phoneNumber, redirectTo } from '@/validation';

describe('Unit tests on field validation', () => {
  describe('phone number', () => {
    it('requires a value', () => {
      const error = phoneNumber.validate(undefined).error!;
      expect(error).not.toBeUndefined();
      const {
        details: [{ message }],
      } = error;
      expect(message).toBe('"value" is required');
    });

    it('should not accept an empty phone number', () => {
      const error = phoneNumber.validate('').error!;
      expect(error).not.toBeUndefined();
      const {
        details: [{ message }],
      } = error;
      expect(message).toBe('"value" is not allowed to be empty');
    });

    it('should not accept an invalid phone number', () => {
      const error = phoneNumber.validate('+12345678').error!;
      expect(error).not.toBeUndefined();
      const {
        details: [{ message }],
      } = error;
      expect(message).toBe(
        '"value" failed custom validation because invalid phone number'
      );
    });

    it('should transform a phone number to its international format', () => {
      const { error, value } = phoneNumber.validate('+33140506070')!;
      expect(error).toBeUndefined();
      expect(value).toEqual('+33140506070');
    });

    it("should not accept phone numbers that don't start with '+'", () => {
      const error = phoneNumber.validate('33140506070').error!;
      expect(error).not.toBeUndefined();
      const {
        details: [{ message }],
      } = error;
      expect(message).toBe(
        '"value" failed custom validation because invalid phone number'
      );
    });

    it("should accept phone numbers starts with '00'", () => {
      const { error, value } = phoneNumber.validate('0033140506070')!;
      expect(error).toBeUndefined();
      expect(value).toEqual('+33140506070');
    });
  });

  describe('redirections', () => {
    const clientUrl = 'https://nhost.io';
    const diffDomainUrl = 'https://myotherdomain.com'
    const host = 'host.com'
    const allowedRedirectUrls = `https://*-nhost.vercel.app,${diffDomainUrl},https://*.${host}`;

    beforeAll(async () => {
      await request.post('/change-env').send({
        AUTH_CLIENT_URL: clientUrl,
        AUTH_ACCESS_CONTROL_ALLOWED_REDIRECT_URLS: allowedRedirectUrls,
      });
    });

    it('should validate any url when the client url is not set', async () => {
      await request.post('/change-env').send({
        AUTH_CLIENT_URL: '',
      });
      expect(
        redirectTo.validate('https://www.google.com/subpath?key=value').value
      ).toEqual('https://www.google.com/subpath?key=value');

      await request.post('/change-env').send({
        AUTH_CLIENT_URL: clientUrl,
      });
    });

    it('should reject an invalid url', () => {
      expect(redirectTo.validate('not-an-url').error).toBeObject();
    });

    it('should validate if it matches the client url', () => {
      expect(redirectTo.validate(clientUrl).error).toBeUndefined();
      expect(redirectTo.validate(`${clientUrl}/path`).error).toBeUndefined();
      expect(redirectTo.validate(`${clientUrl}/key=value`).error).toBeUndefined();
      //expect(redirectTo.validate(`${clientUrl}#key=value`).error).toBeUndefined();
    });

    it('should validate if it matches one of the valid allowed redirect urls', () => {
      expect(redirectTo.validate(`${diffDomainUrl}`).error).toBeUndefined();
      expect(redirectTo.validate(`${diffDomainUrl}/path`).error).toBeUndefined();
      expect(redirectTo.validate(`${diffDomainUrl}/key=value`).error).toBeUndefined();
      expect(redirectTo.validate(`${diffDomainUrl}#key=value`).error).toBeUndefined();
    });

    it('should work with wildcards in the allowed redirect urls', () => {
      expect(redirectTo.validate(`https://subdomain.${host}`).error).toBeUndefined()
      expect(redirectTo.validate(`https://subdomain.${host}/path`).error).toBeUndefined()
      expect(redirectTo.validate(`https://subdomain.${host}#key=value`).error).toBeUndefined()
      expect(redirectTo.validate(`https://subdomain.${host}?key=value`).error).toBeUndefined()

      expect(redirectTo.validate(`https://docs-ger4gr-nhost.vercel.app`).error).toBeUndefined()
    });

    it('localhost', () => {
      expect(redirectTo.validate(`http://localhost:3000`).error).toBeUndefined()
    });

    it('should reject url with the wrong path', () => {
      expect(redirectTo.validate(`https://${host}/wrong`).error).toBeObject();
    });

    it('should reject sub-subdomains if no wildcard', () => {
      expect(
        redirectTo.validate(`https://subdomain.${diffDomainUrl}`).error
      ).toBeObject();
    });

    it('should reject shadowing domains', async () => {
      expect(
        redirectTo.validate(`https://protocol.com.example.com`).error
      ).toBeObject();

      expect(
        redirectTo.validate(`https://wwwaprotocol.com`).error
      ).toBeObject();
    });
  });
});
