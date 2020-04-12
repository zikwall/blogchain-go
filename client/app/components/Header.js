import { useState } from 'react';
import { Input, Menu, Dropdown, Image } from 'semantic-ui-react'
import { connect } from 'react-redux';
import { authenticate, reauthenticate, deauthenticate } from '../redux/actions';

const Header = ({ isAuthenticated }) => {
    const [ activeItem, setActiveItem ] = useState('home');

    const onItemClick = (e, { name }) => {
        setActiveItem(name);
    };

    const trigger = (
        <span>
            <Image avatar src={'https://avatars1.githubusercontent.com/u/23422968?s=460&u=2b4cedc533303cca1599e8785c1f33462251ae9a&v=4'} />
            Andrey Ka
        </span>
    );

    return (
        <Menu secondary>
            <Menu.Item>
                <img src={'/images/bc_300.png'} />
            </Menu.Item>
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
            <Menu.Menu position='right'>
                <Menu.Item>
                    <Input icon='search' placeholder='Search...' />
                </Menu.Item>
                {
                    isAuthenticated &&
                    <Menu.Item>
                        <Dropdown trigger={trigger} pointing='top right'>
                            <Dropdown.Menu pointing secondary>
                                <Dropdown.Item text={
                                    <span>
                                    Signed in as <strong>Andrey Ka</strong>
                                </span>
                                } disabled/>
                                <Dropdown.Item text='Your Profile' selected/>
                                <Dropdown.Item text='Your Stars'/>
                                <Dropdown.Item text='Explore'/>
                                <Dropdown.Item text='Integrations'/>
                                <Dropdown.Item text='Help'/>
                                <Dropdown.Item text='Settings'/>
                                <Dropdown.Divider />
                                <Dropdown.Item text='Sign Out' onClick={() => alert('Log out!')}/>
                            </Dropdown.Menu>
                        </Dropdown>
                    </Menu.Item>
                }
            </Menu.Menu>
        </Menu>
    )
};

const mapStateToProps = (state) => (
    { isAuthenticated: !!state.authentication.token }
);

export default connect(mapStateToProps, { authenticate, reauthenticate, deauthenticate })(Header);
