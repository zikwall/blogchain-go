import { createRef, useState } from "react";
import {
    Segment,
    Ref,
    Grid,
    Container,
    Icon,
    Sticky,
    Menu,
    Label,
    Header,
    Image,
    Button,
    List,
    Divider
} from 'semantic-ui-react';
import { useRouter } from 'next/router';
import { CommonLayout } from "../app/layouts";
import { MenuItemLink } from "../app/components";

const TabMenu = () => (
    <Menu pointing secondary>
        <MenuItemLink href="/" name="Статьи" />
        <MenuItemLink href="/news" name="Новости" />
        <MenuItemLink href="/authors" name="Авторы" />
        <MenuItemLink href="/companies" name="Компании" />
    </Menu>
);

const TagBar = ({ tags }) => {
    return (
        <div style={{ paddingBottom: '20px' }}>
            {tags.map((v, k) => (
                <Label key={k} as='a' horizontal>
                    { v }
                </Label>
            ))}
        </div>
    )
};

const ImageWrap = ({ src }) => (
    <div style={{ paddingBottom: '10px' }}>
        <Image src={src} centered />
    </div>
);

const Publisher = ({ name, time, avatar }) => (
    <>
        <Label as='a' basic image>
            <img src={ avatar } />
            { name }
        </Label>
        <Label as='a' basic>
            { time }
        </Label>
    </>
);

const Article = ({ title, text, image, tags, labels, publisher }) => {
    return (
        <Segment>

            <Publisher
                name={publisher.author}
                time={publisher.time}
                avatar={publisher.avatar}
            />

            <Header as='h2'>
                <a href="/" style={{
                    textDecoration: 'none',
                    color: 'rgba(0,0,0,.87)'
                }}>
                    { title }
                </a>
            </Header>

            {
                tags &&
                <TagBar tags={tags} />
            }

            {
                image &&
                <ImageWrap src={image} />
            }

            { text }

            <Label basic pointing>
                <Icon name='lightning' /> { labels.ratings }
            </Label>
            <Label basic pointing>
                <Icon name='eye' /> { labels.views }
            </Label>
            <Label basic pointing>
                <Icon name='bookmark' /> { labels.bookmarks }
            </Label>
            <Label basic color='blue' pointing>
                Comments
                <Label.Detail>{ labels.comments }</Label.Detail>
            </Label>
        </Segment>
    )
};

export default function Index() {
    const [ activeItem, setActiveItem ] = useState('home');

    const contextRef = createRef();
    const router = useRouter();

    const onItemClick = (e, { name }) => {
        setActiveItem(name);
        router.push('/post/13');
    };

    return (
        <CommonLayout>
            <Container>
                <Grid>
                    <Ref innerRef={contextRef}>
                        <Grid.Row columns={2}>
                            <Grid.Column width={12}>
                                <TabMenu />

                                <Article
                                    publisher={{
                                        author: 'LonelyHeart',
                                        time: 'вчера в 19:07',
                                        avatar: "https://habrastorage.org/getpro/habr/avatars/99b/81e/fff/99b81efff158c12f3c45e55c8a5d6abc.jpg"
                                    }}
                                    tags={[
                                        "Разработка под Arduino",
                                        "Периферия",
                                        "DIY или Сделай сам"
                                    ]}
                                    title="Велотренажер #Самоизоляция или как угомонить ребенка на карантине"
                                    text={
                                        <>
                                            <p>
                                                Весь мир героически борется с «… заразой коронавирусной» (Путин В.В.) Большинство стран закрывают свои границы, своих граждан закрывают на карантин, вводят комендантский час. Вот и Россию не обошла эта гадость стороной.
                                            </p>
                                            <p>
                                                В сложившейся ситуации с пандемией SARS-CoV-2 (COVID-19) все мы с вами сейчас должны находиться на карантине самоизоляции.
                                            </p>
                                            <p>
                                                Поэтому вопрос о том, как найти активное развлечение для детей запертых в четырёх стенах стоит как никогда остро. Надо ещё и постараться чтоб эти четыре стены остались, по-возможности, в целости и сохранности.
                                            </p>
                                        </>
                                    }
                                    labels={{
                                        ratings: 10,
                                        views: 23,
                                        bookmarks: 5,
                                        comments: 214
                                    }}
                                />
                                <Article
                                    publisher={{
                                        author: 'SashulyaK',
                                        time: 'сегодня в 13:08',
                                        avatar: "https://habrastorage.org/getpro/habr/avatars/791/217/d31/791217d314e7458aef0f63497e212538.png"
                                    }}
                                    tags={[
                                        "Блог компании Онлайн-кинотеатр ivi",
                                        "Работа с видео",
                                        "Алгоритмы",
                                    ]}
                                    title="Как мы научились делить видео на сцены с помощью хитрой математики"
                                    text={
                                        <>
                                            <p>
                                                За 10 лет существования ivi мы собрали базу из 90000 видео разной длины, размера и качества. Каждую неделю появляются сотни новых. У нас есть гигабайты метаданных, которые полезны для рекомендаций, упрощают навигацию по сервису и настройку рекламы. Но извлекать информацию непосредственно из видео мы начали только два года назад.
                                            </p>
                                            <p>
                                                В этой статье я расскажу, как мы разбираем фильмы на структурные элементы и зачем нам это нужно. В конце есть ссылка на репозиторий Github с кодом алгоритмов и примерами.
                                            </p>
                                        </>
                                    }
                                    image="/images/tmp/k4psugpjprxnpen_xkfykzemrqg.png"
                                    labels={{
                                        ratings: 10,
                                        views: 23,
                                        bookmarks: 5,
                                        comments: 214
                                    }}
                                />
                                <Article
                                    publisher={{
                                        author: 'baragol',
                                        time: 'сегодня в 14:56',
                                        avatar: "https://habrastorage.org/getpro/habr/avatars/59c/ee7/489/59cee748971fb3479686a26b3b93908a.jpg"
                                    }}
                                    tags={[
                                        'Блог компании Хабр,',
                                        'Здоровье гика'
                                    ]}
                                    title="Проверка: насколько «здоровая» у вас поза, когда работаете за компьютером?"
                                    text={
                                        <>
                                            <p>
                                                Ну ладно, домашними рабочими местами похвастались, теперь давайте о действительно важных вещах — о здоровье. Около 8 часов в сутки мы проводим сидя в одной, часто неудобной позе. Некоторые оригиналы — стоя, — но сути это не меняет. Если не уследить, то через несколько лет незаметно испортится осанка, начнутся проблемы с позвоночником, а оттуда головные боли и прочие неприятности. Мы не претендуем на научность, но пообщавшись с несколькими врачами и перечитав кучу релевантных постов тут на Хабре, соорудили короткий тест, который хотя бы в первом приближении покажет, насколько «здоровая» у вас поза для работы.
                                            </p>
                                        </>
                                    }
                                    image="/images/tmp/d9hit92csqfpdamska7jsmv0pry.jpeg"
                                    labels={{
                                        ratings: 10,
                                        views: 23,
                                        bookmarks: 5,
                                        comments: 214
                                    }}
                                />

                                <BottomPagination />
                                <MostReading />
                            </Grid.Column>
                            <Grid.Column width={4}>
                                <Sticky context={contextRef} offset={30}>
                                    <Menu pointing secondary vertical fluid>
                                        {Flows.map((v, k) => (
                                            <MenuItemLink
                                                key={k}
                                                href={v.href}
                                            >
                                                { v.title }
                                                <Label basic color='green'>{ v.count }</Label>
                                            </MenuItemLink>
                                        ))}
                                    </Menu>
                                </Sticky>
                            </Grid.Column>
                        </Grid.Row>
                    </Ref>
                </Grid>
            </Container>
        </CommonLayout>
    );
}

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

const BottomPagination = () => (
    <div style={{ display: 'flex', flexDirection: 'row', justifyContent: 'space-between', }}>
        <div>
            <Button.Group>
                <Button labelPosition='left' icon='left chevron' content='Туда' />
                <Button labelPosition='right' icon='right chevron' content='Сюда' />
            </Button.Group>
        </div>
        <div>
            <Button.Group>
                <Button icon>
                    1
                </Button>
                <Button icon>
                    2
                </Button>
                <Button icon>
                    3
                </Button>
                <Button icon>
                    4
                </Button>
                <Button icon>
                    5
                </Button>
            </Button.Group>{' '}

            <Button.Group>
                <Button icon>
                    10
                </Button>
                <Button icon>
                    20
                </Button>
                <Button icon>
                    30
                </Button>
            </Button.Group>
        </div>
    </div>
);

const MostReadingLabels = ({ ups, views, bookmarks, comments }) => (
    <>
        <Label basic>
            <Icon name='arrow up' /> { ups }
        </Label>
        <Label basic>
            <Icon name='eye' /> { views }
        </Label>
        <Label basic>
            <Icon name='bookmark' /> { bookmarks }
        </Label>
        <Label basic as="a">
            <Icon name='comment' /> { comments }
        </Label>
    </>
);

const MostReadingItem = ({ labels, title, href }) => {
    return (
        <>
            <Header as="h4">
                <a href={href} style={{
                    textDecoration: 'none',
                    color: 'rgba(0,0,0,.87)'
                }}>
                    { title }
                </a>
            </Header>
            <MostReadingLabels
                ups={labels.ups}
                views={labels.views}
                bookmarks={labels.bookmarks}
                comments={labels.comments}
            />
            <Divider />
        </>
    )
};

const MostReading = () => {
    return (
        <>
            <Divider />

            <div style={{
                backgroundColor: '#f7f7f7',
                padding: '15px',
                borderRadius: '5px'
            }}>
                <Header as="h3" color='grey'>
                    Самое читаемое
                </Header>

                <Divider />

                <div>
                    <Button.Group>
                        <Button basic color='blue'>Сутки</Button>
                        <Button basic color='blue'>Неделя</Button>
                        <Button basic color='blue'>Месяц</Button>
                    </Button.Group>
                </div>

                <MostReadingItem
                    title="Apple выпустила новый iPhone SE по цене от 40 000 ₽"
                    href="/"
                    labels={{
                        ups: 10,
                        views: 25,
                        bookmarks: 6,
                        comments: 23
                    }}
                />
                <MostReadingItem
                    title="Девочки сидят дома: регистрация новых вебкам-моделей выросла на 37—69%"
                    href="/"
                    labels={{
                        ups: 2,
                        views: 14,
                        bookmarks: 2,
                        comments: 6
                    }}
                />
                <MostReadingItem
                    title="Коллеги: и не друг, и не враг, а как?"
                    href="/"
                    labels={{
                        ups: 23,
                        views: 45,
                        bookmarks: 13,
                        comments: 55
                    }}
                />
                <MostReadingItem
                    title="Ликбез по респираторам. Помогает ли респиратор от заражения вирусом. Обзор 11 респираторов"
                    href="/"
                    labels={{
                        ups: 7,
                        views: 30,
                        bookmarks: 6,
                        comments: 16
                    }}
                />
                <MostReadingItem
                    title="Велотренажер #Самоизоляция или как угомонить ребенка на карантине"
                    href="/"
                    labels={{
                        ups: 35,
                        views: 342,
                        bookmarks: 25,
                        comments: 234
                    }}
                />
            </div>
        </>
    )
};
