import { combineReducers } from 'redux'
import authentication, { getUser, getToken } from './auth';

const rootReducer = combineReducers({
    authentication
});

export default rootReducer;

export {
    getToken, getUser
}
