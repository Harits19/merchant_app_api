# Majoo Test Case Backend

Nomor 1 :
- H
Untuk struktur ERD masih bisa ditingkatkan atau dibenahi dengan menghilangkan Foreign Key merchant_id di Table Transactions dikarenakan dari Table Merchant sudah memiliki hubungan one to many pada Tabel Outlet dan Tabel Outlet sendiri juga memiliki hubungan one to many pada Table Transaction. Oleh karena itu, merchant_id dapat dihilangkan dikarenakan Table Merchant dapat terhubung secara tidak langsung ke Table Transaction melalui Table Outlet.
- I
Berikut merupakan DML untuk mengambil data dari tabel transaksi yang dihubungkan oleh tabel outlet, merchant, dan user

    select
	    MIN(M.merchant_name) as merchant_name,
	    MIN(O.outlet_name) as outlet_name ,
	    SUM(T.bill_total) as total_omzet,
	    DATE_FORMAT(T.updated_at, '%Y-%m-%d') as date
    
    from
	    users as U
	    join merchants as M on
	    M.user_id = U.id
	    join outlets as O on
	    M.id = O.merchant_id
	    join transactions as T on
	    T.outlet_id = O.id
    
    where
	    U.id = ?
	    and DATE_FORMAT(T.updated_at, '%Y-%m') = ?
    
    group by
	    date

