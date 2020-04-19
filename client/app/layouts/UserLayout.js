import ProtectedLayout from "./ProtectedLayout";
import Head from "next/head";
import { Button, Container, Grid, Header, Icon, Image, Label, Menu, Ref, Sticky} from "semantic-ui-react";
import { createRef } from 'react';
import { MenuItemLink } from "../components";

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
        <MenuItemLink href="/u/[username]" as="/u/zikwall" name="Обзор" />
        <MenuItemLink href="/u/[username]/all" as="/u/zikwall/all" name="Все статьи" />
        <MenuItemLink href="/u/[username]/stars" as="/u/zikwall/stars" name="Звезды" />
        <MenuItemLink href="/u/[username]/followers" as="/u/zikwall/followers" name="Подписчики" />
        <MenuItemLink href="/u/[username]/followings" as="/u/zikwall/followings" name="Подписки" />
    </Menu>
);

const UserLayout = ({ username, children }) => {
    const contextRef = createRef();

    return (
        <ProtectedLayout>
            <Head>
                <title>ZikWall</title>
            </Head>
            <Container>
                <Grid>
                    <Ref innerRef={contextRef}>
                        <Grid.Row columns={2}>
                            <Grid.Column width={5}>
                                <Sticky context={contextRef} offset={30}>
                                    <Sidebar />
                                </Sticky>
                            </Grid.Column>
                            <Grid.Column width={11}>
                                <TabBar />
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