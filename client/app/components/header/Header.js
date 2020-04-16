import { useState } from 'react';
import { Input, Menu, Image } from 'semantic-ui-react'
import { connect } from 'react-redux';
import { authenticate, reauthenticate } from '../../redux/actions';
import ProfileMenu from './ProfileMenu';
import { useRouter } from "next/router";
import { MenuItemLink } from "../ui/Active";

const Header = ({ isAuthenticated }) => {
    const [ activeItem, setActiveItem ] = useState('home');
    const router = useRouter();

    const onItemClick = (e, { name }) => {
        setActiveItem(name);
    };

    return (
        <Menu secondary>
            <Menu.Item>
                <img src={'/images/bc_300.png'} />
            </Menu.Item>

            <MenuItemLink name="Home" href="/" />
            <MenuItemLink name="Editor" href="/editor" />

            <Menu.Item
                name='messages'
                active={activeItem === 'messages'}
                onClick={() => {
                    router.push('/editor')
                }}
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

                <ProfileMenu />
            </Menu.Menu>
        </Menu>
    )
};

const mapStateToProps = (state) => ({
    isAuthenticated: !!state.authentication.token
});

export default connect(mapStateToProps, { authenticate, reauthenticate })(Header);
