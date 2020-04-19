import UserLayout from "../../../../app/layouts/UserLayout";
import { apiFetch } from "../../../../app/services/api";

const Index = ({ user }) => {
    return (
        <UserLayout user={user}>
            Звезды
        </UserLayout>
    )
};

Index.getInitialProps = async ({ query, req, res }) => {
    const { username } = query;
    const response = await apiFetch(`/api/v1/profile/${username}`, {}, req);

    if (response.status === 100) {
        res.statusCode = 404;
        res.end('Not found');
        return;
    }

    return { user: response.user }
};

export default Index;