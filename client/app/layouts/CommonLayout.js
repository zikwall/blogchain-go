import { Header } from "../components";
import { Container, Menu, Icon, Button, Image } from 'semantic-ui-react';
import { MenuItemLink } from "../components";

const CommonLayout = ({ children }) => (
    <>
        <Header />

        <main className="app root-content">
            { children }
        </main>

        <Container fluid>
            <div style={{
                display: 'flex',
                flexDirection: 'row',
                justifyContent: 'space-between'
            }}>
                <div>
                    <Menu text>
                        <Menu.Item header>
                            <span style={{ paddingRight: '5px' }}>Crafted with</span> <Icon name='heart outline' /> by zikwall
                        </Menu.Item>
                        <MenuItemLink name="О Сайте" href="/" />
                        <MenuItemLink name="Служба поддержки" href="/" />
                        <MenuItemLink name="Мобильные приложения" href="/" />
                    </Menu>
                </div>
                <div style={{
                    display: 'flex',
                    flexDirection: 'row',
                    alignItems: 'center',
                }}>
                    <div style={{
                        display: 'flex',
                        flexDirection: 'row',
                        paddingRight: '20px'
                    }}>
                        <div style={{ paddingRight: '5px' }}>
                            <Image src="/images/google.png" size="tiny" />
                        </div>
                        <div style={{ paddingRight: '5px' }}>
                            <Image src="/images/ios.png" size="tiny" />
                        </div>
                    </div>
                    <div>
                        <Button circular color='facebook' icon='facebook' />
                        <Button circular color='twitter' icon='twitter' />
                        <Button circular color='linkedin' icon='linkedin' />
                        <Button circular color='google plus' icon='google plus' />
                    </div>
                </div>
            </div>
        </Container>
    </>
);

export default CommonLayout;
