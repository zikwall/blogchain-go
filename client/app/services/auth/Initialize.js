import Router from 'next/router';
import { reauthenticate } from '../../redux/actions';
import { Cookie } from '../../help';
import { USER_KEY, SESSION_TOKEN_KEY } from "../../constants";

// checks if the page is being loaded on the server, and if so, get auth token from the cookie:
export default (ctx) => {
    if(ctx.isServer) {
        if(ctx.req.headers.cookie) {
            ctx.store.dispatch(
                reauthenticate(
                    Cookie.getCookie(SESSION_TOKEN_KEY, ctx.req),
                    JSON.parse(Cookie.getCookie(USER_KEY, ctx.req))
                )
            );
        }
    } else {
        const token = ctx.store.getState().authentication.token;

        if(token && (ctx.pathname === '/login' || ctx.pathname === '/register')) {
            setTimeout(function() {
                Router.push('/');
            }, 0);
        }
    }
};
