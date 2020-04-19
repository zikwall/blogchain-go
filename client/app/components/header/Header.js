import { useState } from 'react';
import { Input, Menu, Image } from 'semantic-ui-react'
import { connect } from 'react-redux';
import { authenticate, reauthenticate } from '../../redux/actions';
import ProfileMenu from './ProfileMenu';
import { useRouter } from "next/router";
import { MenuItemLink } from "../ui/Active";

const Header = ({ isAuthenticated }) => {
    return (
        <Menu secondary>
            <Menu.Item>
                <img src={'/images/bc_300.png'} />
            </Menu.Item>

            <MenuItemLink name="Моя лента" href="/" as="/"/>
            <MenuItemLink href="/editor1" as="/editor1" name="Все потоки" />
            <MenuItemLink href="/editor2" as="/editor2" name="Как стать автором" />
            <MenuItemLink href="/editor" as="/editor" name="Новая публикаця!" />

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
