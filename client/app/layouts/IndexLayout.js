import { Container, Grid, Label, Menu, Ref, Sticky } from "semantic-ui-react";
import CommonLayout from "./CommonLayout";
import { MenuItemLink } from "../components";
import { createRef } from "react";

const Flows = [
    {title: 'Разработка', count: '+55', href: '/flows/develop'},
    {title: 'Научоп', count: '+6', href: '/'},
    {title: 'Разработка', count: '+7', href: '/flows/develop'},
    {title: 'Администрирвоание', count: '+22', href: '/flows/develop'},
    {title: 'Дизайн', count: '+4', href: '/flows/develop'},
    {title: 'Разработка', count: '+16', href: '/flows/develop'},
    {title: 'Менеджмент', count: '+13', href: '/flows/develop'},
    {title: 'Маркетинг', count: '+1', href: '/flows/develop'},
];

const TabMenu = () => (
    <Menu pointing secondary>
        <MenuItemLink href="/" name="Статьи" as="/" />
        <MenuItemLink href="/news" name="Новости" as="/news" />
        <MenuItemLink href="/authors" name="Авторы" as="/authors" />
        <MenuItemLink href="/companies" name="Компании" as="companies" />
    </Menu>
);

const IndexLayout = ({ children }) => {
    const contextRef = createRef();

    return (
        <CommonLayout>
            <Container>
                <Grid>
                    <Ref innerRef={contextRef}>
                        <Grid.Row columns={2}>
                            <Grid.Column width={12}>
                                <TabMenu />
                                { children }
                            </Grid.Column>
                            <Grid.Column width={4}>
                                <Sticky context={contextRef} offset={30}>
                                    <Menu pointing secondary vertical fluid>
                                        {Flows.map((v, k) => (
                                            <MenuItemLink
                                                key={k}
                                                href={v.href}
                                                as={v.href}
                                                name={
                                                    <>
                                                        { v.title }
                                                        <Label basic color='green'>{ v.count }</Label>
                                                    </>
                                                }
                                            />
                                        ))}
                                    </Menu>
                                </Sticky>
                            </Grid.Column>
                        </Grid.Row>
                    </Ref>
                </Grid>
            </Container>
        </CommonLayout>
    )
};

export default IndexLayout;