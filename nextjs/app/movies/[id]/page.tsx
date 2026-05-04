type Props = {
    params: Promise<{ id: string }>
}

async function getMovies(id: string) {
    const res = await fetch(`https://api.themoviedb.org/3/movie/${id}`, {
        headers: { Authorization: `Bearer ${process.env.TMDB_TOKEN}`, },
    })
    if (!res.ok) throw new Error('Movie not found')
    return res.json()
}
export default async function MoviePage({ params }: Props) {
    const { id } = await params

    const movie = await getMovies(id)

    return (
        <div>
            <h1>{movie.title}</h1>
            <p>{movie.overview}</p>
            <p>Released: {movie.release_date}</p>
        </div>
    )
}