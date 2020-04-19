import { Editor } from "../app/components";
import { ProtectedLayout } from "../app/layouts";
import { Container, Form, Message, Button } from "semantic-ui-react";
import { createRef, useState } from 'react';
import { Content } from "../app/services";

const EditorPage = () => {
    const contextRef = createRef();
    const [ content, setContent ] = useState('<div><span style="font-size: 18px;">Quill Rich Text Editor</span>'+
        '</div><div><br></div><div>Quill is a free, <a href="https://githu'+
        'b.com/quilljs/quill/">open source</a> WYSIWYG editor built for th'+
        'e modern web. With its <a href="http://quilljs.com/docs/modules/"'+
        '>extensible architecture</a> and a <a href="http://quilljs.com/do'+
        'cs/api/">expressive API</a> you can completely customize it to fu'+
        'lfill your needs. Some built in features include:</div><div><br><'+
        '/div><ul><li>Fast and lightweight</li><li>Semantic markup</li><li'+
        '>Standardized HTML between browsers</li><li>Cross browser support'+
        ' including Chrome, Firefox, Safari, and IE 9+</li></ul><div><br><'+
        '/div><div><span style="font-size: 18px;">Downloads</span></div><d'+
        'iv><br></div><ul><li><a href="https://quilljs.com">Quill.js</a>, '+
        'the free, open source WYSIWYG editor</li><li><a href="https://zen'+
        'oamaro.github.io/react-quill">React-quill</a>, a React component '+
        'that wraps Quill.js</li></ul>'+
        '<code><pre>fn hello() -> Option<u32> {\n' +
        '  Some(1)\n' +
        '}</pre></code>');
    const [ title, setTitle ] = useState("");

    const onSubmit = async () => {
        const { status, message, content_id } = await Content.CreateContent(2, title, content);

        if (status === false) {
            console.log('NOT OK!!!');
            return
        }
    };

    return (
        <ProtectedLayout>
            <Container>
                <Form success>
                    <Form.Group>
                        <Form.Input onChange={(e, { name, value }) => setTitle(value)} placeholder='Как думаете назвать свое твоерени?' width={14}/>
                        <Button primary floated='right' onClick={onSubmit}>Отправить!</Button>
                    </Form.Group>

                    <Message
                        success
                        header='Form Completed'
                        content="You're all signed up for the newsletter"
                    />
                    <div style={{ paddingBottom: '10px' }} />
                </Form>
                <Editor initialValue={content}
                onChange={setContent}
                />
            </Container>
        </ProtectedLayout>
    )
};

export default EditorPage;
