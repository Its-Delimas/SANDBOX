import Link from 'next/link'

async function getPopularMovies() {
    const res = await fetch('https://api.themoviedb.org/3/movie/popular', {
        headers: {
            Authorization: `Bearer ${process.env.TMDB_TOKEN}`,
        },
    })

    if (!res.ok) throw new Error('Failed to fetch movies')

    return res.json()
}

export default async function MoviesPage() {
    const data = await getPopularMovies()
    const movies = data.results

    return (
        <div>
            <h1>Popular Movies</h1>
            <ul>
                {movies.map((movie: any) => (
                    <li key={movie.id}>
                        <Link href={`/movies/${movie.id}`}>
                            {movie.title}
                        </Link>
                    </li>
                ))}
            </ul>
        </div>
    )
}