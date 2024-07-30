import { error } from '@sveltejs/kit';
/** @type {import('./$types').PageLoad} */

export async function load({ params }: { params: { id2: string, id: string } }) {
    const res = await fetch(`https://latte.cf.aayus.me/${params.id}/${params.id2}`,{
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
        }
    });
    const data = await res.json();
    if (res.ok)
        return { title: data.title, content: data.content };
    error(404, 'Not found');
}
