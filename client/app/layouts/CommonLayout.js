import { Header } from "../components";

const CommonLayout = ({ children }) => (
    <>
        <Header />

        <main className="app root-content">
            { children }
        </main>
    </>
);

export default CommonLayout;
