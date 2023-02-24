import { Router } from 'express';

import { asyncWrapper as aw } from '@/utils';
import { bodyValidator } from '@/validation';

import {
  signInEmailPasswordHandlerLegacy,
  signInEmailPasswordSchema,
} from './email-password';

const router = Router();

/**
 * POST /auth/login
 * @summary Authenticate with email + password
 * @param {SignInEmailPasswordSchema} request.body.required
 * @return {SessionPayload} 200 - Signed in successfully - application/json
 * @return {InvalidRequestError} 400 - The payload is invalid - application/json
 * @return {UnauthorizedError} 401 - Invalid email or password, or user is not verified - application/json
 * @return {DisabledEndpointError} 404 - The feature is not activated - application/json
 * @tags Authentication
 */
router.post(
  '/auth/login',
  bodyValidator(signInEmailPasswordSchema),
  aw(signInEmailPasswordHandlerLegacy)
);

const authRouter = router;
export { authRouter };
