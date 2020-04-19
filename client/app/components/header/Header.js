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

            <MenuItemLink name="Моя лента" href="/" as="/"/>
            <MenuItemLink href="/editor" as="/editor" name="Все потоки" />
            <MenuItemLink href="/editor" as="/editor" name="Как стать автором" />

            <Menu.Menu position='right'>
                <Menu.Item>
                    <Input icon='search' placeholder='Поиск...' />
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
