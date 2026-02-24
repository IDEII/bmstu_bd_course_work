import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, map, Observable, throwError } from 'rxjs';
import { Fighter } from './fighter';

@Injectable({
  providedIn: 'root'
})
export class FighterService {

  private apiUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) { }

  getFighters(): Observable<Fighter[]> {
    return this.http.get(`${this.apiUrl}/fighters`).pipe(
      map((data: any) => {
        let fighterlist = data["fighterlist"];
        return fighterlist.map(function(fighter: any): Fighter {
          return new Fighter(
            fighter.user_id,
            fighter.club_id,
            fighter.id,
            fighter.name,
            fighter.description,
            fighter.country,
            new Date(fighter.birthday),
            fighter.rating,
            fighter.category,
          );
        });
      })
    );
  }

  getFighter(id: string): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/fighter/${id}`);
  }

  private handleError(error: any): Observable<never> {
    console.error('An error occurred:', error);
    return throwError(() => new Error(error.message || 'Server error'));
  }

  addFighter(Fighter: any, user_id: string): Observable<any> {
    const newFighter = new HttpHeaders().set("Accept", "application/json");
    return this.http.post<any>(`${this.apiUrl}/add_fighter/${user_id}`, JSON.stringify(Fighter), {headers: newFighter}).pipe(catchError(this.handleError));
  }


  updateFighter(id: string, fighter: any): Observable<any> {
    const updateFighter = new HttpHeaders().set("Accept", "application/json");

    return this.http.post(`${this.apiUrl}/fighter/${id}/update`, JSON.stringify(fighter), {headers: updateFighter}).pipe(
      catchError(error => {
        console.error('Ошибка при обновлении бойца:', error);
        throw error;
      })
    );
  }

  deleteFighter(id: string): Observable<any> {
    return this.http.get(`${this.apiUrl}/fighter/${id}/delete`);
  }

}
