import { Request, Response, Router } from "express";
import { manageProviderStrategy } from "./utils";
import { PROVIDERS } from "@/config";
import { OAuth2Client } from "google-auth-library";
import { getSignInResponse } from "@/utils/tokens";
import { asyncWrapper as aw } from "@/helpers";
import { User } from "@/types";

const options = PROVIDERS.google;

const client = new OAuth2Client(options?.clientID)

// const transformProfile = ({
//       id,
//       emails,
//       displayName,
//       photos,
//     }: Profile): {
//       id: string;
//       email?: string;
//       displayName: string;
//       avatarUrl?: string;
//     } => ({
//       id,
//       email: emails?.[0]?.value,
//       displayName: displayName,
//       avatarUrl: photos?.[0]?.value || getGravatarUrl(emails?.[0].value),
//     })
const transformProfile = ({ email, displayName, name, picture, sub }: {
  sub: string;
  displayName: string
  name: string,
  picture: string;
  email: string;
  email_verified: boolean;
  locale: string;
}): { id: string; email: string; displayName: string; avatarUrl: string } => ({
  id: sub,
  email: email as string,
  displayName: displayName ?? name,
  avatarUrl: picture,
})

// NEW
const googleOneTap = async (req: Request, res: Response): Promise<unknown> => {
  const { credential } = req.body;
  
  const ticket = await client.verifyIdToken({
    idToken: credential as string,
  })

  const payload = ticket.getPayload()

  const account = (await manageProviderStrategy('google', transformProfile as any)(req as any, '', '', payload as any, (err: any, user: any) => {
    return user;
  })) as unknown as User

  // COPY: login
  if (!account) {
    return res.boom.unauthorized('Invalid code');
  }
  
  // if (account.disabled) {
  //   return res.boom.badRequest('User is disabled');
  // }
  
  const signInResponse = await getSignInResponse({
    userId: account.id,
    checkMFA: false,
  });
  
  return res.send(signInResponse);
};
  

// router.post(
//   '/signin/passwordless/sms',
//   createValidator().body(signInPasswordlessSmsSchema),
//   aw(signInPasswordlessSmsHandler)
// );
export default (router: Router) => router.post('/google-one-tap', aw(googleOneTap))




// export default (router: Router) => router.post('/google-one-tap', asyncWrapper(googleOneTap))
