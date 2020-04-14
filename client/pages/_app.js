import { Provider } from "react-redux";
import App from "next/app";
import withRedux from "next-redux-wrapper";
import { MainLayout } from "../app/layouts";
import { makeStore } from "../app/redux/Store";
import { Initialize } from '../app/services/auth';

// assets
import 'semantic-ui-css/semantic.min.css'
import '../app/assets/custom.css';
import 'react-quill/dist/quill.snow.css';

class MyApp extends App {
    static async getInitialProps({ Component, ctx }) {
        Initialize(ctx);

        return {
            pageProps: {
                ...(Component.getInitialProps
                    ? await Component.getInitialProps(ctx)
                    : {})
            }
        };
    }

    render() {
        const { Component, pageProps, store } = this.props;

        return (
            <Provider store={ store }>
                <MainLayout>
                    <Component {...pageProps} />
                </MainLayout>
            </Provider>
        )
    }
}

export default withRedux(makeStore)(MyApp);
