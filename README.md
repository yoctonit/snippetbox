Great question — yes, you can install and run MySQL on Windows without registering it as a Windows service, so you can start and stop it manually when needed.

Here’s the clean, step-by-step method 👇

🧰 1. Download the ZIP (not Installer) version

Go to the official MySQL downloads page:
👉 https://dev.mysql.com/downloads/mysql/
Choose Windows (x86, 64-bit), ZIP Archive — not the MSI Installer.

Extract it to a folder, e.g.:
C:\mysql

⚙️ 2. Create a basic configuration file

Inside C:\mysql, create a file named my.ini with contents like:
[mysqld]
basedir=C:\mysql
datadir=C:\mysql\data
port=3306
sql_mode=NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES

📁 3. Initialize the data directory

Open Command Prompt as Administrator, then:
cd C:\mysql\bin
mysqld --initialize --console

This will:

Create the data folder.

Generate system tables.

Print a temporary root password (copy it — you'll need it).

If you want an empty root password, use:
mysqld --initialize-insecure

▶️ 4. Start MySQL manually (no service)

From the same folder:
mysqld --console
Keep that console open — MySQL is now running.

If you want it to run in the background (still manually started):
start mysqld

⏹️ 5. Stop MySQL manually

In another Command Prompt window:
mysqladmin -u root -p shutdown

Enter the password when prompted.
🧑‍💻 6. Connect to the server

While it’s running:
mysql -u root -p

⚡ Optional shortcuts

You can make batch files for convenience:

start_mysql.bat
@echo off
cd C:\mysql\bin
start mysqld

stop_mysql.bat
@echo off
cd C:\mysql\bin
mysqladmin -u root -p shutdown

✅ Summary
| Action         | Command                              |
| -------------- | ------------------------------------ |
| Initialize     | `mysqld --initialize --console`      |
| Start manually | `mysqld --console` or `start mysqld` |
| Stop manually  | `mysqladmin -u root -p shutdown`     |
| Connect        | `mysql -u root -p`                   |


Perfect 👌 — let’s go over how to set up MySQL as a self-contained instance inside your Go project directory, so you can start and stop it manually — great for local dev work (like your snippetbox project).

🗂️ 1. Folder structure (example)

Let’s assume your project folder looks like this:
C:\Projects\snippetbox\
│
├── mysql\
│   ├── bin\
│   ├── data\
│   ├── my.ini
│
├── app\
├── go.mod
└── main.go

So MySQL lives inside your project under mysql/.

⚙️ 2. Download MySQL ZIP (no installer)

Download the ZIP archive for Windows from
👉 https://dev.mysql.com/downloads/mysql/
[mysqld]
basedir=C:/Projects/snippetbox/mysql
datadir=C:/Projects/snippetbox/mysql/data
port=3306
sql_mode=NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES
skip-networking=0
skip-grant-tables=0

(Use forward slashes / — MySQL prefers that on Windows.)

If you plan to run multiple MySQL instances, you can use a different port (e.g. 3307).

⚡ 4. Initialize the data directory

Open Command Prompt (as Admin or not, both work), and run:
cd C:\Projects\snippetbox\mysql\bin
mysqld --defaults-file=../my.ini --initialize-insecure
mysqld --defaults-file="C:\Data\bin\mysql\my.ini" --initialize-insecure --basedir="C:\Data\bin\mysql" --datadir="C:\Data\bin\mysql\data"

This will:

Create the data directory.

Initialize an empty root password.

▶️ 5. Start MySQL manually

Start it from your project folder:
cd C:\Projects\snippetbox\mysql\bin
mysqld --defaults-file=../my.ini --console
mysqld --defaults-file="C:\Data\bin\mysql\my.ini" --basedir="C:\Data\bin\mysql" --datadir="C:\Data\bin\mysql\data" --console


Keep this window open while the server runs.

You can also run it in background:
start mysqld --defaults-file=../my.ini

⏹️ 6. Stop MySQL manually

In a new terminal:
cd C:\Projects\snippetbox\mysql\bin
mysqladmin --defaults-file=../my.ini -u root shutdown
mysqladmin --defaults-file="C:\Data\bin\mysql\my.ini" -u root shutdown

💻 7. Connect to the local instance

While it’s running:
cd C:\Projects\snippetbox\mysql\bin
mysql -u root

You’re now in the MySQL shell.

🔧 8. (Optional) Add helper scripts

Inside your project root, create:

start_mysql.bat
@echo off
cd mysql\bin
start mysqld --defaults-file=../my.ini
echo MySQL started.

stop_mysql.bat
@echo off
cd mysql\bin
mysqladmin --defaults-file=../my.ini -u root shutdown
echo MySQL stopped.

You can now double-click those from Explorer or run:
start_mysql
and
stop_mysql

✅ 9. Use in your Go app

Your Go app can now connect using:
dsn := "root@tcp(127.0.0.1:3306)/snippetbox?parseTime=true"
db, err := sql.Open("mysql", dsn)

🧩 Benefits of this setup

No system-wide MySQL service.

Fully portable — works from USB, ZIP, or repo.

You can include your own pre-initialized schema under version control (e.g., in mysql/init.sql).

Great for dev/demo environments.

ALTER USER 'root'@'localhost' IDENTIFIED BY 'Password!';
FLUSH PRIVILEGES;
