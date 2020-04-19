import {apiFetch} from "../api";

export const CreateContent = (user_id, title, content) => {
    return apiFetch("/api/v1/content/add", {
        method: 'POST',
        body: JSON.stringify({
            user_id: user_id,
            title: title,
            content: content
        })
    }).then((res) => {
        if (res.status === 100) {
            return {
                status: false,
                message: res.message,
                content_id: 0
            }
        }

        return {
            status: true,
            message: res.message,
            content_id: res.content_id
        }
    })
};

export const GetContent = (id) => {
    return apiFetch(`/api/v1/content/${id}`).then((res) => {
        if (res.status === 100) {
            return {
                status: false,
                content: "",
                title: "",
                user: null
            }
        }

        return {
            status: true,
            content: res.content,
            title: res.title,
            user: res.user
        }
    })
};