import { Dropdown, Menu } from "semantic-ui-react";
import { withRouter } from 'next/router';
import Link from "next/link";

const MenuItemLink = withRouter(({ href, name, as, router }) => {
    return <Menu.Item
        active={(router.asPath === href || router.asPath === as)}
        onClick={() => {
            router.push(as)
        }}
    >
        <Link href={href} as={as}>
            <>
                { name }
            </>
        </Link>
    </Menu.Item>
});

const DropdownItemLink = withRouter(({ name, href, as, router }) => {
    return <Dropdown.Item 
        text={name}
        selected={(router.asPath === href || router.asPath === as)}
        onClick={() => {
            router.push(href)
        }}
    >
        <Link href={href} as={as}>
            <>
                { name }
            </>
        </Link>
    </Dropdown.Item>
});

export {
    MenuItemLink, DropdownItemLink
}
