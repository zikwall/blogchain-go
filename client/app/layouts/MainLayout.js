import Head from "next/head";
import { Container} from 'semantic-ui-react';

const MainLayout = ({ children }) => {
    return (
        <Container fluid className="root-container">
            <Head>
                <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
                <link rel="icon" href="/favicon.ico"/>
            </Head>

            { children }
        </Container>
    )
};

export default MainLayout;
