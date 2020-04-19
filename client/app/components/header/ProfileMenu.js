import { useState } from 'react';
import { Dropdown, Image, Menu } from "semantic-ui-react";
import { connect } from 'react-redux';
import { deauthenticate } from '../../redux/actions';
import { bindActionCreators } from "redux";
import AuthItem from "./AuthItem";
import QuickLogin from "./QuickLogin";
import CloseWrapper from "../close/CloseWrapper";
import { DropdownItemLink } from '../ui/Active';

const ProfileMenu = ({ isAuthenticated, logout, user }) => {
    const [ isDropped, setIsDropped ] = useState(false);
    const onOutsideOrEscape = () => {
        setIsDropped(false);
    };

    if (!isAuthenticated) {
        return (
            <>
                <Menu.Item
                    onClick={() => {
                        setIsDropped(!isDropped)
                    }}>

                    <AuthItem />
                </Menu.Item>

                <CloseWrapper onEscapeOutside={ onOutsideOrEscape }>
                    <QuickLogin visible={isDropped} />
                </CloseWrapper>
            </>
        )
    }

    const trigger = (
        <span>
            <Image avatar src={'https://avatars1.githubusercontent.com/u/23422968?s=460&u=2b4cedc533303cca1599e8785c1f33462251ae9a&v=4'} />
            { user.profile.name }
        </span>
    );

    return (
        <Menu.Item>
            <Dropdown trigger={trigger} pointing='top right'>
                <Dropdown.Menu pointing secondary>
                    <Dropdown.Item text={
                        <span>
                         Авторизированы как <strong>{ user.username }</strong>
                    </span>
                    } disabled/>

                    <DropdownItemLink name='Мой профиль' href={`/u/${user.username}`} />
                    <DropdownItemLink name='Новый пост' href='/editor' />
                    <DropdownItemLink name='Мои звезды' href={`/u/${user.username}/stars`} />
                    <DropdownItemLink name='Публикации' href={`/u/${user.username}/all`} />
                    <DropdownItemLink name='Диалоги' href='/dialogs' />
                    <DropdownItemLink name='Закладки' href='/bookmarks' />
                    <DropdownItemLink name='Помощь' href='/help' />
                    <DropdownItemLink name='Настройки' href='/settings' />

                    <Dropdown.Divider />

                    <Dropdown.Item text='Выйти' onClick={() => logout()}/>
                </Dropdown.Menu>
            </Dropdown>
        </Menu.Item>
    )
};

const mapStateToProps = (state) => ({
    isAuthenticated: !!state.authentication.token,
    user: state.authentication.user
});

const mapDispatchToProps = dispatch => bindActionCreators({
    logout: deauthenticate
}, dispatch);


export default connect(mapStateToProps, mapDispatchToProps)(ProfileMenu);
