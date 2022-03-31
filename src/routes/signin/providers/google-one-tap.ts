export const a = 1
// import { Request, Response, Router } from "express";
// import { PROVIDERS } from "@/config";
// import { asyncWrapper } from "../../../../../hasura-backend-plus/src/shared/helpers";
// import { OAuth2Client } from "google-auth-library";
// import { setRefreshToken } from "../../../../../hasura-backend-plus/src/shared/cookies";
// import { Session, User } from "@/types";
// import { createHasuraJwt } from "../../../../../hasura-backend-plus/src/shared/jwt";
// import { manageProviderStrategy } from "@/routes/signin/providers/utils";
//
//
// // export default (router: Router): void => {
// //   const options = PROVIDERS.google;
// //
// //   initProvider(
// //     router,
// //     'google',
// //     Strategy,
// //     {
// //       scope: PROVIDERS.google?.scope,
// //       prompt: 'consent',
// //       accessType: 'offline',
// //     },
// //     (req, res, next) => {
// //       if (!PROVIDERS.google) {
// //         return res.boom.notImplemented(`Google sign-in is not enabled`);
// //       } else if (!options?.clientID || !options?.clientSecret) {
// //         throw new Error(`Missing environment variables for Google OAuth`);
// //       } else {
// //         return next();
// //       }
// //     }
// //   );
// // };
//
//
//
// const client = new OAuth2Client(PROVIDERS.google?.clientID)
//
// const transformProfile = ({ email, displayName, picture, sub }: {
//   sub: string;
//   displayName: string;
//   picture: string;
//   email: string;
//   email_verified: boolean;
//   locale: string;
// }): { id: string; email: string; displayName: string; avatarUrl: string } => ({
//   id: sub,
//   email: email as string,
//   displayName,
//   avatarUrl: picture,
// })
//
// // NEW
// const googleOneTap = async (req: Request, res: Response): Promise<void> => {
//   const { body } = req;
//
//   // default to true
//   const useCookie = typeof body.cookie !== 'undefined' ? body.cookie : true
//
//   const ticket = await client.verifyIdToken({
//     idToken: body.credential as string,
//   })
//
//   const payload = ticket.getPayload()
//
//   const account = (await manageProviderStrategy('google', transformProfile as any)(req, '', '', payload as any, (err: any, user: any) => {
//     return user;
//   })) as unknown as User
//
//   // COPY: login
//   const refresh_token = await setRefreshToken(res, account.id, useCookie)
//
//   const jwt_token = createHasuraJwt(account)
//   const jwt_expires_in = newJwtExpiry
//
//   const session: Session = { accessToken: jwt_token, accessTokenExpiresIn: jwt_expires_in, user: account.user }
//   if (!useCookie) session.refreshToken = refresh_token
//
//   res.send(session)
//
// };
//
// export default (router: Router) => router.post('/google-one-tap', asyncWrapper(googleOneTap))
