import Head from "next/head";
import { Button, Card, Image, Form, Header, Grid } from 'semantic-ui-react';

const Login = () => {
    return (
        <>
            <Head>
                <title>Blog | Auth</title>
            </Head>

            <Grid textAlign='center' style={{ height: '75vh' }} verticalAlign='middle'>
                <Grid.Column style={{ maxWidth: 450 }}>
                    <Header as='h2' textAlign='center'>
                        <Image src={'/images/bc_300.png'} />

                        <span style={{marginRight: 20, marginLeft: 10, verticalAlign: 'middle'}}>
                            Log-in to your account
                        </span>
                    </Header>
                    <Card fluid>
                        <Card.Content>
                            <Form size='large'>

                                <Form.Input fluid icon='user' iconPosition='left' placeholder='E-mail address' />
                                <Form.Input
                                    fluid
                                    icon='lock'
                                    iconPosition='left'
                                    placeholder='Password'
                                    type='password'
                                />
                                <Button fluid>
                                    Login
                                </Button>
                            </Form>
                        </Card.Content>
                        <Card.Content extra>
                            New to us? <a href='#'>Sign Up</a>
                        </Card.Content>
                    </Card>
                </Grid.Column>
            </Grid>
        </>
    )
};

export default Login;
