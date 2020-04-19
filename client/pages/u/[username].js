import { ProtectedLayout} from "../../app/layouts";
import Head from "next/dist/next-server/lib/head";
import {
    Container,
    Grid,
    Ref,
    Image,
    Header,
    Button,
    Label,
    Icon,
    Menu,
    Segment
} from "semantic-ui-react";
import { createRef } from 'react';
import { MenuItemLink } from "../../app/components";
import { LabelBar } from "../../app/components/Article";

const Profile = () => {
    const contextRef = createRef();

    return (
        <ProtectedLayout>
            <Head>
                <title>Как мы научились делить видео на сцены с помощью хитрой математики</title>
            </Head>
            <Container>
                <Grid>
                    <Ref innerRef={contextRef}>
                        <Grid.Row columns={2}>
                            <Grid.Column width={5}>
                                <Sidebar />
                            </Grid.Column>
                            <Grid.Column width={11}>
                                <TabBar />
                                <Content />
                            </Grid.Column>
                        </Grid.Row>
                    </Ref>
                </Grid>
            </Container>
        </ProtectedLayout>
    )
};

const Sidebar = () => (
    <div style={{ width: '300px' }}>
        <div style={{
            border: '1px solid rgba(0,0,0,.1)',
        }}>
            <Image
                src='https://avatars1.githubusercontent.com/u/23422968?s=460&u=2b4cedc533303cca1599e8785c1f33462251ae9a&v=4'
                size='medium'
            />
            <div style={{
                padding: '10px',
            }}>
                Study...
            </div>
        </div>

        <div style={{ paddingTop: '10px', paddingBottom: '10px'}}>
            <Header as="h1">
                <div style={{
                    display: 'flex',
                    flexDirection: 'column'
                }}>
                    <span>Andrey Kapitonov</span>
                </div>
            </Header>
            <span style={{
                fontHeight: 300,
                fontSize: '24px',
                lineHeight: '14px',
            }}>zikwall</span>
        </div>
        <div style={{ paddingTop: '10px', paddingBottom: '10px' }}>
            <Button animated='fade' basic fluid>
                <Button.Content visible>Редактировать профиль</Button.Content>
                <Button.Content hidden>Приступить</Button.Content>
            </Button>
        </div>

        <div style={{ paddingTop: '10px', paddingBottom: '10px' }}>
            #PHP, #Go, #JS, #React, #ReactNative - full stack developer
            TODO: #Rust
        </div>
        <div style={{ paddingTop: '10px', paddingBottom: '10px' }}>
            <Label basic>
                <Icon name='send' /> andrey.kapitonov@gmail.com
            </Label>
            <div style={{ paddingTop: '5px' }} />
            <Label basic>
                <Icon name='map marker alternate' /> Russian, Moscow
            </Label>
        </div>
    </div>
);

const TabBar = () => (
    <Menu pointing secondary>
        <MenuItemLink href="/profile">
            Обзор
        </MenuItemLink>
        <MenuItemLink href="/profile/all">
            Все статьи
        </MenuItemLink>
        <MenuItemLink href="/profile/stars">
            Лайки
        </MenuItemLink>
        <MenuItemLink href="/profile/followers">
            Подписчики
        </MenuItemLink>
        <MenuItemLink href="/profile/followings">
            Подписки
        </MenuItemLink>
    </Menu>
);

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
