import App from "next/app";
import { MainLayout } from "../app/layouts";

// assets
import 'semantic-ui-css/semantic.min.css'
import '../app/assets/custom.css';

class MyApp extends App {
    static async getInitialProps({ Component, ctx }) {
        return {
            pageProps: {
                ...(Component.getInitialProps
                    ? await Component.getInitialProps(ctx)
                    : {})
            }
        };
    }

    render() {
        const { Component, pageProps } = this.props;

        return (
            <MainLayout>
                <Component {...pageProps} />
            </MainLayout>
        )
    }
}

export default MyApp;
