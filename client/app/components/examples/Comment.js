import { Button, Comment, Form, Header } from 'semantic-ui-react'

const CommentExampleThreaded = () => (
    <Comment.Group threaded className={'root-comments'}>
        <Header as='h3' dividing>
            Комментарии
        </Header>

        <Comment>
            <Comment.Avatar as='a' src='https://habrastorage.org/getpro/habr/avatars/791/217/d31/791217d314e7458aef0f63497e212538.png' />
            <Comment.Content>
                <Comment.Author as='a'>SashulyaK</Comment.Author>
                <Comment.Metadata>
                    <span>вчера в 13:59</span>
                </Comment.Metadata>
                <Comment.Text>Вы правы, мы изначально этим пользовались. Но разница между соседними фреймами даёт ложные срабатывания в динамичных сценах. А ещё она не детектит переходы, когда шоты плавно перетекают друг в друга фейдами. В итоге пришлось перейти на нейронки.</Comment.Text>
                <Comment.Actions>
                    <a>Ответить</a>
                </Comment.Actions>
            </Comment.Content>
        </Comment>

        <Comment>
            <Comment.Avatar as='a' src='https://habrastorage.org/getpro/habr/avatars/93f/965/c06/93f965c063c9c1910aac6271e3686cc9.jpg' />
            <Comment.Content>
                <Comment.Author as='a'>GeneD88</Comment.Author>
                <Comment.Metadata>
                    <span>вчера в 23:56</span>
                </Comment.Metadata>
                <Comment.Text>
                    <p>А как подобный алгоритм справляется с такими фильмами, где присутствует сплит-экран, например как The Sisters Де Пальмы (2 экрана), Timecode Фиггиса (4 экрана), Charly Нельсона и тп?</p>
                </Comment.Text>
                <Comment.Actions>
                    <a>Ответить</a>
                </Comment.Actions>
            </Comment.Content>

            <Comment.Group>
                <Comment>
                    <Comment.Avatar as='a' src='https://habrastorage.org/getpro/habr/avatars/791/217/d31/791217d314e7458aef0f63497e212538.png' />
                    <Comment.Content>
                        <Comment.Author as='a'>SashulyaK</Comment.Author>
                        <Comment.Metadata>
                            <span>Только что</span>
                        </Comment.Metadata>
                        <Comment.Text>Хороший воспрос) Если использовать только фичи видеоряда, то, скорее всего, не сработает. Но если добавить звук (или использовать только звук), это может решить проблему.</Comment.Text>
                        <Comment.Actions>
                            <a>Ответить</a>
                        </Comment.Actions>
                    </Comment.Content>
                </Comment>

                <Comment>
                    <Comment.Avatar as='a' src='https://habrastorage.org/getpro/habr/avatars/279/fd4/4f1/279fd44f1aab1c22791c701977182439.png' />
                    <Comment.Content>
                        <Comment.Author as='a'>marsermd</Comment.Author>
                        <Comment.Metadata>
                            <span>вчера в 23:39</span>
                        </Comment.Metadata>
                        <Comment.Text>Хабр здорового человека, где снова пишут про разработку</Comment.Text>
                        <Comment.Actions>
                            <a>Ответить</a>
                        </Comment.Actions>
                    </Comment.Content>
                </Comment>
            </Comment.Group>
        </Comment>

        <Comment>
            <Comment.Avatar as='a' src='https://habrastorage.org/getpro/habr/avatars/279/fd4/4f1/279fd44f1aab1c22791c701977182439.png' />
            <Comment.Content>
                <Comment.Author as='a'>marsermd</Comment.Author>
                <Comment.Metadata>
                    <span>вчера в 23:39</span>
                </Comment.Metadata>
                <Comment.Text>Какая приятная статья! Давненько такого на хабре не читал. Спасибо большое!</Comment.Text>
                <Comment.Actions>
                    <a>Ответить</a>
                </Comment.Actions>
            </Comment.Content>
        </Comment>

        <Form reply>
            <Form.TextArea />
            <Button content='Отправить' labelPosition='left' icon='edit' primary />
        </Form>
    </Comment.Group>
)

export default CommentExampleThreaded
