import Head from "next/head";
import { Container} from 'semantic-ui-react';

const MainLayout = ({ children }) => {
    return (
        <Container fluid className="root-container">
            { children }
        </Container>
    )
};

export default MainLayout;
