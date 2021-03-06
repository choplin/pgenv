package main

func ExampleListCommand_pretty() {
	app := makeTestEnv()
	app.Run([]string{"pgenv", "list", "-f", "pretty"})
	// Output:
	// +-------------+------------------+
	// |    NAME     |     VERSION      |
	// +-------------+------------------+
	// | 9.3.9-debug | PostgreSQL 9.3.9 |
	// | 9.4.4       | PostgreSQL 9.4.4 |
	// +-------------+------------------+
}

func ExampleListCommand_prettyDetail() {
	app := makeTestEnv()
	app.Run([]string{"pgenv", "list", "-f", "pretty", "-d"})
	// Output:
	// +-------------+------------------+---------------+------------------------------------------+--------------------------------------------+--------------------------------------------+
	// |    NAME     |     VERSION      | GIT REFERENCE |                   HASH                   |                    PATH                    |             CONFIGURE OPTIONS              |
	// +-------------+------------------+---------------+------------------------------------------+--------------------------------------------+--------------------------------------------+
	// | 9.3.9-debug | PostgreSQL 9.3.9 | REL9_3_9      | 553e576e05b50f9faffbd3dd721e44fc3746898d | /home/postgres/.pgenv/versions/9.3.9-debug | --prefix                                   |
	// |             |                  |               |                                          |                                            | /home/postgres/.pgenv/versions/9.3.9-debug |
	// |             |                  |               |                                          |                                            | --enable-debug --enable-cassert            |
	// | 9.4.4       | PostgreSQL 9.4.4 | REL9_4_4      | 7c055f3ec3bd338a1ebb8c73cff3d01df626471e | /home/postgres/.pgenv/versions/9.4.4       | --prefix                                   |
	// |             |                  |               |                                          |                                            | /home/postgres/.pgenv/versions/9.4.4       |
	// +-------------+------------------+---------------+------------------------------------------+--------------------------------------------+--------------------------------------------+
}

func ExampleListCommand_plain() {
	app := makeTestEnv()
	app.Run([]string{"pgenv", "list", "-f", "plain"})
	// Output:
	// Name	Version
	// 9.3.9-debug	PostgreSQL 9.3.9
	// 9.4.4	PostgreSQL 9.4.4
}

func ExampleListCommand_plainDetail() {
	app := makeTestEnv()
	app.Run([]string{"pgenv", "list", "-f", "plain", "-d"})
	// Output:
	// Name	Version	Git Reference	Hash	Path	Configure Options
	// 9.3.9-debug	PostgreSQL 9.3.9	REL9_3_9	553e576e05b50f9faffbd3dd721e44fc3746898d	/home/postgres/.pgenv/versions/9.3.9-debug	--prefix /home/postgres/.pgenv/versions/9.3.9-debug --enable-debug --enable-cassert
	// 9.4.4	PostgreSQL 9.4.4	REL9_4_4	7c055f3ec3bd338a1ebb8c73cff3d01df626471e	/home/postgres/.pgenv/versions/9.4.4	--prefix /home/postgres/.pgenv/versions/9.4.4
}
func ExampleListCommand_json() {
	app := makeTestEnv()
	app.Run([]string{"pgenv", "list", "-f", "json"})
	// Output: [{"name":"9.3.9-debug","version":"PostgreSQL 9.3.9"},{"name":"9.4.4","version":"PostgreSQL 9.4.4"}]
}

func ExampleListCommand_jsonDetail() {
	app := makeTestEnv()
	app.Run([]string{"pgenv", "list", "-f", "json", "-d"})
	// Output: [{"name":"9.3.9-debug","version":"PostgreSQL 9.3.9","git-ref":"REL9_3_9","hash":"553e576e05b50f9faffbd3dd721e44fc3746898d","path":"/home/postgres/.pgenv/versions/9.3.9-debug","configureOptions":["--prefix","/home/postgres/.pgenv/versions/9.3.9-debug","--enable-debug","--enable-cassert"]},{"name":"9.4.4","version":"PostgreSQL 9.4.4","git-ref":"REL9_4_4","hash":"7c055f3ec3bd338a1ebb8c73cff3d01df626471e","path":"/home/postgres/.pgenv/versions/9.4.4","configureOptions":["--prefix","/home/postgres/.pgenv/versions/9.4.4"]}]
}
