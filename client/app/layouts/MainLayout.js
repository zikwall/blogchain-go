import Head from "next/head";
import { Header as MenuHeader} from "../components";
import { Container} from 'semantic-ui-react';


const MainLayout = ({ children }) => {
    return (
        <Container fluid className="root-container">
            <Head>
                <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
                <link rel="icon" href="/favicon.ico"/>
            </Head>

            <MenuHeader />

            <main className="app root-content">
                { children }
            </main>
        </Container>
    )
};

export default MainLayout;
