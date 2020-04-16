import Head from "next/head";
import { Container} from 'semantic-ui-react';

const MainLayout = ({ children }) => {
    return (
        <>
            <Head>
                <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
                <link rel="icon" href="/favicon.ico"/>
                <link href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/styles/monokai-sublime.min.css" rel="stylesheet" />
                <script src="//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.18.1/highlight.min.js"></script>
                <script src="/js/hls.js" />
                <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/katex@0.11.1/dist/katex.min.css"
                      integrity="sha384-zB1R0rpPzHqg7Kpt0Aljp8JPLqbXI3bhnPWROx27a9N0Ll6ZP/+DiW/UqRcLbRjq"
                      crossOrigin="anonymous" />
                <script defer src="https://cdn.jsdelivr.net/npm/katex@0.11.1/dist/katex.min.js"
                        integrity="sha384-y23I5Q6l+B6vatafAwxRu/0oK/79VlbSz7Q9aiSZUvyWYIYsd+qj+o24G5ZU2zJz"
                        crossOrigin="anonymous">

                </script>
            </Head>
            <Container fluid className="root-container">
                { children }
            </Container>
        </>
    )
};

export default MainLayout;
