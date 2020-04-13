import { useState, createRef } from 'react';
import { useRouter } from 'next/router';
import {
    Container,
    Grid,
    Header,
    Icon,
    Menu,
    Message,
    Ref,
    Sticky
} from "semantic-ui-react";
import CommentExampleThreaded from "../../app/components/examples/Comment";
import { CommonLayout } from "../../app/layouts";

const Post = () => {
    const [ activeItem, setActiveItem ] = useState('home');

    const contextRef = createRef();
    const router = useRouter();

    const onItemClick = (e, { name }) => {
        setActiveItem(name);
        router.push('/');
    };

    return (
        <CommonLayout>
            <Container>
                <Grid>
                    <Ref innerRef={contextRef}>
                        <Grid.Row columns={2}>
                            <Grid.Column width={13}>
                                <Container text fluid>
                                    <Header as='h2'>Header</Header>
                                    <p>
                                        Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo
                                        ligula eget dolor. Aenean massa strong. Cum sociis natoque penatibus et
                                        magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis,
                                        ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa
                                        quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget,
                                        arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo.
                                        Nullam dictum felis eu pede link mollis pretium. Integer tincidunt. Cras
                                        dapibus. Vivamus elementum semper nisi. Aenean vulputate eleifend tellus.
                                        Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim.
                                        Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus
                                        viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet.
                                        Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi.
                                    </p>
                                    <Message icon>
                                        <Icon name='circle notched' loading />
                                        <Message.Content>
                                            <Message.Header>Just one second</Message.Header>
                                            We are fetching that content for you.
                                        </Message.Content>
                                    </Message>
                                    <p>
                                        Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo
                                        ligula eget dolor. Aenean massa strong. Cum sociis natoque penatibus et
                                        magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis,
                                        ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa
                                        quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget,
                                        arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo.
                                        Nullam dictum felis eu pede link mollis pretium. Integer tincidunt. Cras
                                        dapibus. Vivamus elementum semper nisi. Aenean vulputate eleifend tellus.
                                        Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim.
                                        Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus
                                        viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet.
                                        Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi.
                                    </p>
                                    <p>
                                        Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo
                                        ligula eget dolor. Aenean massa strong. Cum sociis natoque penatibus et
                                        magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis,
                                        ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa
                                        quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget,
                                        arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo.
                                        Nullam dictum felis eu pede link mollis pretium. Integer tincidunt. Cras
                                        dapibus. Vivamus elementum semper nisi. Aenean vulputate eleifend tellus.
                                        Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim.
                                        Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus
                                        viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet.
                                        Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi.
                                    </p>
                                    <p>
                                        Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo
                                        ligula eget dolor. Aenean massa strong. Cum sociis natoque penatibus et
                                        magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis,
                                        ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa
                                        quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget,
                                        arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo.
                                        Nullam dictum felis eu pede link mollis pretium. Integer tincidunt. Cras
                                        dapibus. Vivamus elementum semper nisi. Aenean vulputate eleifend tellus.
                                        Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim.
                                        Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus
                                        viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet.
                                        Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi.
                                    </p>
                                    <p>
                                        Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo
                                        ligula eget dolor. Aenean massa strong. Cum sociis natoque penatibus et
                                        magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis,
                                        ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa
                                        quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget,
                                        arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo.
                                        Nullam dictum felis eu pede link mollis pretium. Integer tincidunt. Cras
                                        dapibus. Vivamus elementum semper nisi. Aenean vulputate eleifend tellus.
                                        Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim.
                                        Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus
                                        viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet.
                                        Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi.
                                    </p>
                                    <p>
                                        Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo
                                        ligula eget dolor. Aenean massa strong. Cum sociis natoque penatibus et
                                        magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis,
                                        ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa
                                        quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget,
                                        arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo.
                                        Nullam dictum felis eu pede link mollis pretium. Integer tincidunt. Cras
                                        dapibus. Vivamus elementum semper nisi. Aenean vulputate eleifend tellus.
                                        Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim.
                                        Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus
                                        viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet.
                                        Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi.
                                    </p>

                                    <CommentExampleThreaded />
                                </Container>
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
        </CommonLayout>
    )
};

export default Post;
