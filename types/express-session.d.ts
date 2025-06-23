import type { GrantSession } from 'grant';

// Augment express-session with a custom SessionData object
declare module 'express-session' {
  interface SessionData {
    grant: GrantSession;
    redirectTo: string;
  }
}
