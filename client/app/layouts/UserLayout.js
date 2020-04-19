import ProtectedLayout from "./ProtectedLayout";
import Head from "next/head";
import { Button, Container, Grid, Header, Icon, Image, Label, Menu, Ref, Sticky} from "semantic-ui-react";
import { createRef } from 'react';
import { MenuItemLink } from "../components";

const Sidebar = ({ user }) => {
    const avatar = !!user.profile.avatar ? user.profile.avatar : "/images/zebra_pl.jpg";

    return (
        <div style={{width: '300px'}}>
            <div style={{
                border: '1px solid rgba(0,0,0,.1)',
            }}>
                <Image
                    src={avatar}
                    size='medium'
                />

                {
                    !!user.profile.status &&
                    <div style={{
                        padding: '10px',
                    }}>
                        {user.profile.status}
                    </div>
                }

            </div>

            <div style={{ paddingTop: '10px', paddingBottom: '10px' }}>
                <Header as="h1">
                    <div style={{
                        display: 'flex',
                        flexDirection: 'column'
                    }}>
                        <span>{user.profile.name}</span>
                    </div>
                </Header>
                <span style={{
                    fontHeight: 300,
                    fontSize: '24px',
                    lineHeight: '14px',
                }}>{user.username}</span>
            </div>

            <div style={{paddingTop: '10px', paddingBottom: '10px'}}>
                <Button animated='fade' basic fluid>
                    <Button.Content visible>Редактировать профиль</Button.Content>
                    <Button.Content hidden>Приступить</Button.Content>
                </Button>
            </div>

            {
                !!user.profile.description &&
                <div style={{paddingTop: '10px', paddingBottom: '10px'}}>
                    { user.profile.description }
                </div>
            }

            <div style={{paddingTop: '10px', paddingBottom: '10px'}}>
                {
                    !!user.profile.email &&
                    <Label basic>
                        <Icon name='send'/> {user.profile.email}
                    </Label>
                }
                {
                    !!user.profile.location &&
                    <>
                        <div style={{paddingTop: '5px'}}/>
                        <Label basic>
                            <Icon name='map marker alternate'/> Russian, Moscow
                        </Label>
                    </>
                }
            </div>
        </div>
    )
};

const TabBar = ({ user }) => (
    <Menu pointing secondary>
        <MenuItemLink href="/u/[username]" as={`/u/${user.username}`} name="Обзор" />
        <MenuItemLink href="/u/[username]/all" as={`/u/${user.username}/all`} name="Все статьи" />
        <MenuItemLink href="/u/[username]/stars" as={`/u/${user.username}/stars`} name="Звезды" />
        <MenuItemLink href="/u/[username]/followers" as={`/u/${user.username}/followers`} name="Подписчики" />
        <MenuItemLink href="/u/[username]/followings" as={`/u/${user.username}/followings`} name="Подписки" />
    </Menu>
);

const UserLayout = ({ user, children }) => {
    const contextRef = createRef();

    return (
        <ProtectedLayout>
            <Head>
                <title>{ user.profile.name } | Blogchain</title>
            </Head>
            <Container>
                <Grid>
                    <Ref innerRef={contextRef}>
                        <Grid.Row columns={2}>
                            <Grid.Column width={5}>
                                <Sticky context={contextRef} offset={30}>
                                    <Sidebar user={user}/>
                                </Sticky>
                            </Grid.Column>
                            <Grid.Column width={11}>
                                <TabBar user={user} />
                                { children }
                            </Grid.Column>
                        </Grid.Row>
                    </Ref>
                </Grid>
            </Container>
        </ProtectedLayout>
    )
};

export default UserLayout;