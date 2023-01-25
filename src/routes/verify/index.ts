import { Router } from 'express';

import { asyncWrapper as aw } from '@/utils';
import { verifyHandler, verifySchema } from './verify-email';
import { bodyValidator, queryValidator, signInOtpSchema } from '@/validation';
import { verifyPhoneChangeHandler } from './verify-phone';

const router = Router();

// Workaround for Outlook safe links. See: https://github.com/nhost/hasura-auth/issues/189
router.head('/verify', (_, res) => res.sendStatus(200));

// TODO: use VerifySchema in the jsdoc
/**
 * GET /verify
 * @summary Verify tickets created by email verification, email passwordless authentication, or password reset
 * @param {string} ticket.query.required - Ticket generated in the previous actions and sent by email
 * @param {string} type.query.required - name param description - enum:emailVerify,emailConfirmChange,signinPasswordless,passwordReset
 * @param {string} redirectTo.query.required - Redirection link
 * @return {string} 302 - {redirectTo}?refreshToken=${refreshToken}&type=${type}
 * @return {InvalidRequestError} 400 - The payload format is invalid - application/json
 * @tags General
 */
router.get('/verify', queryValidator(verifySchema), aw(verifyHandler));

/**
 * POST /verify/otp
 * @summary Verify phone number change via one-time password code received by SMS
 * @param {SignInOtpSchema} request.body.required
 * @return {string} 200 - User successfully vefiried their new phone number - application/json
 * @return {InvalidRequestError} 400 - The payload is invalid - application/json
 * @return {UnauthorizedError} 401 - Error processing the request - application/json
 * @return {DisabledEndpointError} 404 - The feature is not activated - application/json
 * @security BearerAuth
 * @tags General
 */
router.post(
  '/verify/otp',
  bodyValidator(signInOtpSchema),
  aw(verifyPhoneChangeHandler)
);

const verifyRouter = router;
export { verifyRouter };
