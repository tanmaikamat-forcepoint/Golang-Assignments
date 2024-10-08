class DB{

    void openDB(){

    }
}


class FirebaseDB extends  DB{

}

class SQLDB extends  DB {

}


class DBManager{
    Db database;

    public DBManager(Db db) {
        database=db;
        database.openDb()
    }
    
}