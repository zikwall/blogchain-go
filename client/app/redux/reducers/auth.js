import { AUTHENTICATE, DEAUTHENTICATE, REAUTHENTICATE } from '../types';

const initialState = {
    token: null,
    user: {
        id: 0,
        username: 'Не определен',
        email: 'Не определн',
        profile: {
            name: null,
            public_email: null,
            avatar: null
        }
    },
};

const authentication = (state = initialState, action) => {
    switch(action.type) {
        case AUTHENTICATE:
            return { token: action.token, user: action.user};
        case DEAUTHENTICATE:
            return { token: null };
        default:
            return state;
    }
};

export default authentication;

export const getUser = state => state.authentication.user;
export const getToken = state => state.authentication.token;

