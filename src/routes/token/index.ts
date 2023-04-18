import { Router } from 'express';

import { asyncWrapper as aw } from '@/utils';
import { bodyValidator } from '@/validation';
import {
  createPATHandler,
  createPATSchema,
  tokenHandler,
  tokenSchema,
  verifyPATHandler,
  verifyPATSchema,
} from './token';
import { verifyTokenHandler, verifyTokenSchema } from './verify';

const router = Router();

/**
 * POST /token
 * @summary Refresh the access JWT token
 * @param {TokenSchema} request.body.required
 * @return {SessionModel} 200 - User successfully authenticated - application/json
 * @return {InvalidRequestError} 400 - The payload is invalid - application/json
 * @return {UnauthorizedError} 401 - Unauthenticated user or invalid token - application/json
 * @tags General
 */
router.post('/token', bodyValidator(tokenSchema), aw(tokenHandler));

/**
 * POST /token/verify
 * @summary Veify JWT token
 * @description If request body is not passed the autorization header will be used to be verified
 * @param {VerifyTokenSchema} request.body.optional
 * @return {string} 200 - Valid JWT token - application/json
 * @return {UnauthorizedError} 401 - Unauthenticated user or invalid token - application/json
 * @tags General
 */
router.post(
  '/token/verify',
  bodyValidator(verifyTokenSchema),
  aw(verifyTokenHandler)
);

router.post('/pat', bodyValidator(createPATSchema), aw(createPATHandler));

router.post(
  '/pat/verify',
  bodyValidator(verifyPATSchema),
  aw(verifyPATHandler)
);

const tokenRouter = router;
export { tokenRouter };
