import {Dropdown, Menu} from "semantic-ui-react";
import { useRouter } from 'next/router';

const MenuItemLink = ({ href, ...props }) => {
    const router = useRouter();
    return <Menu.Item
        {...props}
        active={router.pathname === href}
        onClick={() => {
            router.push(href)
        }}
    />
};

const DropdownItemLink = ({ name, href }) => {
    const router = useRouter();
    return <Dropdown.Item 
        text={name}
        selected={router.pathname === href}
        onClick={() => {
            router.push(href)
        }}
    />
};

export {
    MenuItemLink, DropdownItemLink
}
