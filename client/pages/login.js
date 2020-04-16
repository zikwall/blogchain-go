import React, { useEffect, useState } from 'react';
import Head from "next/head";
import { useRouter } from "next/router";
import { Button, Card, Image, Form, Header, Grid, Message } from 'semantic-ui-react';
import { authenticate } from "../app/redux/actions";
import { connect } from "react-redux";
import { bindActionCreators } from 'redux';
import { WithoutHeaderLayout } from "../app/layouts";

const Login = ({ isAuthenticated, auth }) => {
    const [ username, setUsername ] = useState('');
    const [ password, setPassword ] = useState('');
    const [ error, setError ] = useState({
        has: false, message: ''
    });
    const router = useRouter();

    useEffect(() => {
        if (isAuthenticated) {
            router.push('/');
        }

        return () => {}
    }, []);

    const onChangeUsername = (e) => {
        e.preventDefault();

        setUsername(e.target.value);
    };

    const onChangePassword = (e) => {
        e.preventDefault();

        setPassword(e.target.value);
    };

    const onSubmit = async (e) => {
        e.preventDefault();
        const { status, message } = await auth({ username: username, password: password }, 'login');

        if (status === false) {
            setError({
                has: true,
                message: message
            });

            return false;
        }

        if (status === true) {
            router.push('/');
        }
    };

    return (
        <WithoutHeaderLayout>
            <Head>
                <title>Blog | Auth</title>
            </Head>

            <Grid textAlign='center' style={{ height: '85vh' }} verticalAlign='middle'>
                <Grid.Column style={{ maxWidth: 450 }}>
                    <Header as='h2' textAlign='center'>
                        <Image src={'/images/bc_300.png'} />

                        <span style={{marginRight: 20, marginLeft: 10, verticalAlign: 'middle'}}>
                            Log-in to your account
                        </span>
                    </Header>
                    <Card fluid>
                        <Card.Content>
                            {
                                error.has &&
                                <Message
                                    error
                                    header='Action Forbidden'
                                    content={error.message}
                                />
                            }
                            <Form size='large'>

                                <Form.Input
                                    fluid icon='user'
                                    iconPosition='left'
                                    placeholder='E-mail address'
                                    onChange={onChangeUsername}
                                />
                                <Form.Input
                                    fluid
                                    icon='lock'
                                    iconPosition='left'
                                    placeholder='Password'
                                    type='password'
                                    onChange={onChangePassword}
                                />
                                <Button fluid onClick={onSubmit}>
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
        </WithoutHeaderLayout>
    )
};

const mapStateToProps = (state) => (
    { isAuthenticated: !!state.authentication.token }
);

const mapDispatchToProps = dispatch => bindActionCreators({
    auth: authenticate
}, dispatch);

export default connect(mapStateToProps, mapDispatchToProps)(Login);
