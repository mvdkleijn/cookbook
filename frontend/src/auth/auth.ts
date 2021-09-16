import { createOidcAuth, SignInType, LogLevel } from 'vue-oidc-client/vue3';

const loco = window.location;
const appRootUrl = `${loco.protocol}//${loco.host}${process.env.BASE_URL}`;

const idsrvAuth = createOidcAuth(
  'main',
  SignInType.Window,
  appRootUrl,
  {
    authority: process.env.COOKBOOK_AUTHORITY,
    client_id: process.env.COOKBOOK_CLIENTID,
    response_type: 'code',
    redirect_uri: `${appRootUrl}:8081/oidc/callback`,
    scope: 'openid profile email',
    prompt: 'login',
  },
  console,
  LogLevel.Debug,
);

// handle events
idsrvAuth.events.addAccessTokenExpiring(() => console.log('access token expiring'));

idsrvAuth.events.addAccessTokenExpired(() => console.log('access token expired'));

idsrvAuth.events.addSilentRenewError((err) => console.error('silent renew error', err));

idsrvAuth.events.addUserLoaded((user) => console.log('user loaded', user));

idsrvAuth.events.addUserUnloaded(() => console.log('user unloaded'));

idsrvAuth.events.addUserSignedOut(() => console.log('user signed out'));

idsrvAuth.events.addUserSessionChanged(() => console.log('user session changed'));

export default idsrvAuth;
