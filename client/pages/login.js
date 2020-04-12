import React, { useEffect, useState } from 'react';
import Head from "next/head";
import Router from "next/router";
import { Button, Card, Image, Form, Header, Grid } from 'semantic-ui-react';
import { authenticate } from "../app/redux/actions";
import { connect } from "react-redux";
import { bindActionCreators } from 'redux';

const Login = ({ isAuthenticated, auth }) => {
    const [ username, setUsername ] = useState('');
    const [ password, setPassword ] = useState('');
    const [ error, setError ] = useState(false);

    useEffect(() => {
        if (isAuthenticated) {
            Router.push('/');
        }

        return () => {}
    }, []);

    const handleChangeUsername = (e) => {
        e.preventDefault();

        setUsername(e.target.value);
    };

    const handleChangePassword = (e) => {
        e.preventDefault();

        setPassword(e.target.value);
    };

    const handleClickSubmit = (e) => {
        e.preventDefault();
        auth({ username: username, password: password }, 'login');
    };

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

                                <Form.Input
                                    fluid icon='user'
                                    iconPosition='left'
                                    placeholder='E-mail address'
                                    onChange={handleChangeUsername}
                                />
                                <Form.Input
                                    fluid
                                    icon='lock'
                                    iconPosition='left'
                                    placeholder='Password'
                                    type='password'
                                    onChange={handleChangePassword}
                                />
                                <Button fluid onClick={handleClickSubmit}>
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

const mapStateToProps = (state) => (
    { isAuthenticated: !!state.authentication.token }
);

const mapDispatchToProps = dispatch => bindActionCreators({
    auth: authenticate
}, dispatch);

export default connect(mapStateToProps, mapDispatchToProps)(Login);
