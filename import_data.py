import sqlite3
import pandas as pd
from sqlalchemy import create_engine

# GeoPackage is SQLite-based
gpkg_path = '/Users/user/Downloads/grid3-nga-operational-wards-v2.0/GRID3_NGA_operational_wards_v2_0.gpkg'

# Connect to the GeoPackage
conn = sqlite3.connect(gpkg_path)
cursor = conn.cursor()

# List all tables
cursor.execute("SELECT name FROM sqlite_master WHERE type='table';")
tables = cursor.fetchall()
print("Tables in GeoPackage:")
for table in tables:
    print(f"  - {table[0]}")

# Read the main table
if tables:
    # Find the data table (not metadata tables)
    table_name = 'GRID3_NGA_operational_wards_v2_0'
    print(f"\nReading table: {table_name}")
    
    # Get column info
    cursor.execute(f"PRAGMA table_info({table_name});")
    columns = cursor.fetchall()
    print("\nColumns:")
    for col in columns:
        print(f"  - {col[1]} ({col[2]})")
    
    # Read data (exclude geometry column for now)
    df = pd.read_sql_query(f"SELECT OBJECTID, country, iso3, state, statecode, lga, lga_alt_names, ward, ward_alt_names, multipart_count, source, date, area_sqkm FROM {table_name}", conn)
    print(f"\nTotal rows: {len(df)}")
    print("\nFirst few rows:")
    print(df.head())
    
    # Save to CSV for import
    csv_path = '/Users/user/Downloads/fetcher/wards_data.csv'
    df.to_csv(csv_path, index=False)
    print(f"\nData saved to: {csv_path}")
    
    # Import directly to PostgreSQL
    print("\nImporting to PostgreSQL...")
    db_url = "postgresql+psycopg2://postgres:postgres@localhost:5431/grid3_wards"
    engine = create_engine(db_url)
    
    # Import data
    df.to_sql('wards', engine, if_exists='replace', index=False)
    print("Data imported successfully to 'wards' table!")

conn.close()
