import decode from 'jwt-decode';
import { Cookie } from '../../help';
import { SESSION_TOKEN_KEY } from "../../constants";

export default class Session {
    static isGuest = (req) => {
        return !Session.isLogged(req);
    };

    static isLogged = (req) => {
        const token = Cookie.getCookie(SESSION_TOKEN_KEY, req);
        return !!token && !Session.isSessionExpired(token);
    };

    static isSessionExpired = (accessToken) => {
        try {
            const decoded = decode(accessToken);

            return (decoded.exp < Date.now() / 1000);
        } catch (err) {
            console.log('Expired token! Logout...');
            return false;
        }
    };

    static flushSession = () => {
        Cookie.removeCookie(SESSION_TOKEN_KEY);
    };
}
