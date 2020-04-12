import rootReducer from './reducers';
import { createLogger } from 'redux-logger';
import { createStore, applyMiddleware } from 'redux';
import { composeWithDevTools } from 'redux-devtools-extension';
import thunkMiddleware from 'redux-thunk';

const loggerMiddleware = createLogger();

export const makeStore = (initialState, options) => {
    return createStore(rootReducer, initialState,
        composeWithDevTools(applyMiddleware(
            thunkMiddleware,
            loggerMiddleware
        ))
    );
};
