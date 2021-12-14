import { createOidcAuth, SignInType, LogLevel } from 'vue-oidc-client/vue3';

const loco = window.location;
const appRootUrl = `${loco.protocol}//${loco.host}${process.env.BASE_URL}`;

const idsrvAuth = createOidcAuth(
  'main',
  SignInType.Window,
  appRootUrl,
  {
    authority: process.env.VUE_APP_COOKBOOK_AUTHORITY,
    client_id: process.env.VUE_APP_COOKBOOK_CLIENTID,
    response_type: 'code',
    redirect_uri: `${appRootUrl}oidc/callback`,
    scope: 'openid profile email',
    prompt: 'login',
  },
  console,
  LogLevel.Info,
);

// handle events
idsrvAuth.events.addAccessTokenExpiring(() => console.debug('access token expiring'));

idsrvAuth.events.addAccessTokenExpired(() => console.debug('access token expired'));

idsrvAuth.events.addSilentRenewError((err) => console.error('silent renew error', err));

idsrvAuth.events.addUserLoaded((user) => console.debug('user loaded', user));

idsrvAuth.events.addUserUnloaded(() => console.debug('user unloaded'));

idsrvAuth.events.addUserSignedOut(() => console.debug('user signed out'));

idsrvAuth.events.addUserSessionChanged(() => console.debug('user session changed'));

export default idsrvAuth;
