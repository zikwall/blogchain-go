import { createRef, useState } from "react";
import {
    Dimmer,
    Loader,
    Image,
    Segment,
    Ref,
    Grid,
    Container,
    Message,
    Icon,
    Sticky,
    Menu
} from 'semantic-ui-react';
import { useRouter } from 'next/router';

export default function Index() {
    const [ activeItem, setActiveItem ] = useState('home');

    const contextRef = createRef();
    const router = useRouter();

    const onItemClick = (e, { name }) => {
        setActiveItem(name);
        router.push('/post/13');
    };

    return (
        <Container>
            <Grid>
                <Ref innerRef={contextRef}>
                    <Grid.Row columns={2}>
                        <Grid.Column width={13}>
                            <Message icon>
                                <Icon name='circle notched' loading />
                                <Message.Content>
                                    <Message.Header>Just one second</Message.Header>
                                    We are fetching that content for you.
                                </Message.Content>
                            </Message>
                            {[...new Array(15)].map((i, k) => (
                                <Segment key={k}>
                                    <Dimmer active inverted>
                                        <Loader inverted>Loading</Loader>
                                    </Dimmer>

                                    <Image src='https://react.semantic-ui.com/images/wireframe/short-paragraph.png' />
                                </Segment>
                            ))}
                        </Grid.Column>
                        <Grid.Column width={3}>
                            <Sticky context={contextRef} offset={30}>
                                <Menu pointing secondary vertical fluid>
                                    <Menu.Item
                                        name='home'
                                        active={activeItem === 'home'}
                                        onClick={onItemClick}
                                    />
                                    <Menu.Item
                                        name='messages'
                                        active={activeItem === 'messages'}
                                        onClick={onItemClick}
                                    />
                                    <Menu.Item
                                        name='friends'
                                        active={activeItem === 'friends'}
                                        onClick={onItemClick}
                                    />
                                </Menu>
                            </Sticky>
                        </Grid.Column>
                    </Grid.Row>
                </Ref>
            </Grid>
        </Container>
    );
}
