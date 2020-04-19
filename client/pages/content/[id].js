import { Content } from "../../app/services";
import { Container } from 'semantic-ui-react';
import { CommonLayout } from "../../app/layouts";

const ContentPage = ({ content }) => {

    return (
        <CommonLayout>
            <Container>
                <div className="root-dangerous-content-html" dangerouslySetInnerHTML={{ __html: content }} />
            </Container>
        </CommonLayout>
    )
};

ContentPage.getInitialProps = async ({ res, query }) => {
    const { id } = query;
    const { status, content, title, user } = await Content.GetContent(id);

    if (status === false) {
        res.statusCode = 404;
        res.end('Not found');
        return;
    }

    return { content: content }
};

export default ContentPage;