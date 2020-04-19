import {
    Grid,
    Header,
    Segment
} from "semantic-ui-react";
import { LabelBar } from "../../../app/components/Article";
import UserLayout from "../../../app/layouts/UserLayout";
import { apiFetch } from "../../../app/services/api";

const Profile = ({ user }) => {
    return (
        <UserLayout user={user}>
            <Content />
        </UserLayout>
    )
};

// TODO move useEffect UserLayout?
Profile.getInitialProps = async ({ req, query, res }) => {
    const { username } = query;
    const response = await apiFetch(`/api/v1/profile/${username}`, {}, req);

    if (response.status === 100) {
        res.statusCode = 404;
        res.end('Not found');
        return;
    }

    return { user: response.user }
};

const Pinneds = ({ items }) => {
    let groupingItems = [];
    let c = 0;

    for (let group in items) {
        if (group % 2 == 0) {
            c++;
        }

        if (typeof groupingItems[c] === 'undefined') {
            groupingItems[c] = [];
        }

        let item = items[group];
        groupingItems[c].push(
            <Grid.Column width={8}>
                <PinnedItem
                    tags={item.tags}
                    labels={item.labels}
                    text={item.text}
                    title={item.title}
                />
            </Grid.Column>
        );
    }

    return (
        <Grid>
            {
                groupingItems.map(( group, key ) => (
                    <Grid.Row columns={2}>
                        { group }
                    </Grid.Row>
                ))
            }
        </Grid>
    )
};

const PinnedItem = ({ title, labels }) => (
    <Segment>
        <Header as='h5'>
            <a href="/post/13" style={{
                textDecoration: 'none',
                color: 'rgba(0,0,0,.87)'
            }}>
                { title }
            </a>
        </Header>

        <LabelBar
            ratings={labels.ratings}
            bookmarks={labels.bookmarks}
            comments={labels.comments}
            views={labels.views}
        />
    </Segment>
);

const Content = () => (
    <>
        <Pinneds
            items={[
                {
                    title:"Велотренажер #Самоизоляция или как угомонить ребенка на карантине",
                    labels: {
                        ratings: 10,
                        views: 23,
                        bookmarks: 5,
                        comments: 214
                    }
                },
                {
                    title:"Как мы научились делить видео на сцены с помощью хитрой математики",
                    labels: {
                        ratings: 10,
                        views: 23,
                        bookmarks: 5,
                        comments: 214
                    }
                },
                {
                    title:"Как мы научились делить видео на сцены с помощью хитрой математики",
                    labels: {
                        ratings: 10,
                        views: 23,
                        bookmarks: 5,
                        comments: 214
                    }
                },
                {
                    title:"Как мы научились делить видео на сцены с помощью хитрой математики",
                    labels: {
                        ratings: 10,
                        views: 23,
                        bookmarks: 5,
                        comments: 214
                    }
                },
                {
                    title:"Как мы научились делить видео на сцены с помощью хитрой математики",
                    labels: {
                        ratings: 10,
                        views: 23,
                        bookmarks: 5,
                        comments: 214
                    }
                },
                {
                    title:"Как мы научились делить видео на сцены с помощью хитрой математики",
                    labels: {
                        ratings: 10,
                        views: 23,
                        bookmarks: 5,
                        comments: 214
                    }
                }
            ]}
        />
    </>
);

export default Profile;
