import { Heading } from '@chakra-ui/layout'
import type { NextPage } from 'next'
import Head from 'next/head'

const Home: NextPage = () => {
  return (
    <>
      <div className="backdrop">
        <div className="main">
          <header className="header-navbar">
            <Heading as="h1" size="md">
              header
            </Heading>
          </header>
          <footer className="footer-navbar">footer</footer>
        </div>
      </div>
    </>
  )
}

export default Home
