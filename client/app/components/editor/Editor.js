import { useState } from 'react';
import dynamic from 'next/dynamic';

const QuillNoSSRWrapper = dynamic(
    import('react-quill'), { ssr: false, loading: () => <p>Loading ...</p> }
);

const defModules = {
    toolbar: [
        [{ 'header': '1' }, { 'header': '2' }, { 'font': [] }],
        [{ size: [] }],
        ['bold', 'italic', 'underline', 'strike', 'blockquote'],
        [{ 'list': 'ordered' }, { 'list': 'bullet' },
            { 'indent': '-1' }, { 'indent': '+1' }],
        ['link', 'image', 'video', 'formula'],
        ['clean'],
        ['code-block'],
    ],
    syntax: true,
};

const defFormats = [
    'header', 'font', 'size',
    'bold', 'italic', 'underline', 'strike', 'blockquote',
    'list', 'bullet', 'indent',
    'link', 'image', 'video', 'code-block', 'formula'
];

const Editor = ({ modules, formats, onChange, initialValue }) => {
    formats = [ ...defFormats, ...formats ];
    modules = { ...defModules, ...modules };

    const onOverrideChange = (value) => {
        onChange(value);
    };

    return (
        <QuillNoSSRWrapper
            modules={modules}
            formats={formats}
            value={initialValue}
            onChange={onOverrideChange}
        />
    );
};

Editor.defaultProps = {
    formats: [],
    modules: {},
    onChange: () => {},
    initialValue: ''
};

export default Editor;
