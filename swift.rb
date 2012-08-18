#!/usr/bin/env ruby

require 'benchmark'
require 'bundler/setup'
require 'swift-db-postgres'

dbs = {
  swift: Swift::DB::Postgres.new(db: 'swift_test', ssl: {sslmode: 'disable'}),
}

queries = {
  swift: {
    drop:   'drop table if exists users',
    create: 'create table users(id serial primary key, name text, created_at timestamp with time zone)',
    insert: 'insert into users(name, created_at) values (?, ?)',
    select: 'select * from users'
  },
}

rows = 1000
iter = 100

Benchmark.bm(20) do |bm|
  dbs.each do |name, db|
    sql = queries[name]
    db.execute(sql[:drop])
    db.execute(sql[:create])

    bm.report("#{name} insert") do
      rows.times do |n|
        db.execute(sql[:insert], "name #{n}", Time.now)
      end
    end

    bm.report("#{name} select") do
      iter.times do
        db.execute(sql[:select]).entries
      end
    end
  end
end
