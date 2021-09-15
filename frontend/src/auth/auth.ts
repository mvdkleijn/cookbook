// import { User } from 'oidc-client';
import { createOidcAuth, SignInType, LogLevel } from 'vue-oidc-client/vue3';

const loco = window.location;
const appRootUrl = `${loco.protocol}//${loco.host}${process.env.BASE_URL}`;

// const idsrvAuth = createOidcAuth(
//   'main',
//   SignInType.Window,
//   appRootUrl,
//   {
//     authority: process.env.COOKBOOK_AUTHORITY,
//     client_id: process.env.COOKBOOK_CLIENTID,
//     response_type: 'id_token token',
//     redirect_uri: `${appRootUrl}:8080/oidc/callback`,
//     scope: 'openid profile email',
//     prompt: 'login',
//   },
//   console,
//   LogLevel.Debug,
// );

const idsrvAuth = createOidcAuth(
  'main',
  SignInType.Window,
  appRootUrl,
  {
    authority: 'https://hulsbus.eu.auth0.com/',
    client_id: 'z3kP7vq1n8kW3LnImylXJ8YjvFROR8N2',
    response_type: 'id_token token',
    redirect_uri: 'http://localhost:8081/oidc/callback',
    scope: 'openid profile email',
    prompt: 'login',
  },
  console,
  LogLevel.Debug,
);

// // handle events
// idsrvAuth.events.addAccessTokenExpiring(function () {
//   // eslint-disable-next-line no-console
//   console.log('access token expiring');
// });

// idsrvAuth.events.addAccessTokenExpired(function () {
//   // eslint-disable-next-line no-console
//   console.log('access token expired');
// });

// idsrvAuth.events.addSilentRenewError(function (err: Error) {
//   // eslint-disable-next-line no-console
//   console.error('silent renew error', err);
// });

// idsrvAuth.events.addUserLoaded(function (user: User) {
//   // eslint-disable-next-line no-console
//   console.log('user loaded', user);
// });

// idsrvAuth.events.addUserUnloaded(function () {
//   // eslint-disable-next-line no-console
//   console.log('user unloaded');
// });

// idsrvAuth.events.addUserSignedOut(function () {
//   // eslint-disable-next-line no-console
//   console.log('user signed out');
// });

// idsrvAuth.events.addUserSessionChanged(function () {
//   // eslint-disable-next-line no-console
//   console.log('user session changed');
// });

export default idsrvAuth;
